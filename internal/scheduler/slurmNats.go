// Copyright (C) 2023 NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package scheduler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/ClusterCockpit/cc-backend/internal/importer"
	"github.com/ClusterCockpit/cc-backend/internal/repository"
	"github.com/ClusterCockpit/cc-backend/pkg/log"
	"github.com/ClusterCockpit/cc-backend/pkg/schema"
	"github.com/nats-io/nats.go"
)

type SlurmNatsConfig struct {
	URL string `json:"url"`
}

type SlurmNatsScheduler struct {
	url string

	RepositoryMutex sync.Mutex
	JobRepository   *repository.JobRepository
}

type StopJobRequest struct {
	// Stop Time of job as epoch
	StopTime  int64           `json:"stopTime" validate:"required" example:"1649763839"`
	State     schema.JobState `json:"jobState" validate:"required" example:"completed"` // Final job state
	JobId     *int64          `json:"jobId" example:"123000"`                           // Cluster Job ID of job
	Cluster   *string         `json:"cluster" example:"fritz"`                          // Cluster of job
	StartTime *int64          `json:"startTime" example:"1649723812"`                   // Start Time of job as epoch
}

func (sd *SlurmNatsScheduler) startJob(req *schema.JobMeta) {
	fmt.Printf("DEBUG: %+v\n", *req)

	log.Printf("Server Name: %s - Job ID: %v", req.BaseJob.Cluster, req.BaseJob.JobID)
	log.Printf("User: %s - Project: %s", req.BaseJob.User, req.BaseJob.Project)

	if req.State == "" {
		req.State = schema.JobStateRunning
	}
	if err := importer.SanityChecks(&req.BaseJob); err != nil {
		log.Errorf("Sanity checks failed: %s", err.Error())
		return
	}

	// aquire lock to avoid race condition between API calls
	var unlockOnce sync.Once
	sd.RepositoryMutex.Lock()
	defer unlockOnce.Do(sd.RepositoryMutex.Unlock)

	// Check if combination of (job_id, cluster_id, start_time) already exists:
	jobs, err := sd.JobRepository.FindAll(&req.JobID, &req.Cluster, nil)
	if err != nil && err != sql.ErrNoRows {
		log.Errorf("checking for duplicate failed: %s", err.Error())
		return
	} else if err == nil {
		for _, job := range jobs {
			if (req.StartTime - job.StartTimeUnix) < 86400 {
				log.Errorf("a job with that jobId, cluster and startTime already exists: dbid: %d", job.ID)
				return
			}
		}
	}

	id, err := sd.JobRepository.Start(req)
	if err != nil {
		log.Errorf("insert into database failed: %s", err.Error())
		return
	}
	// unlock here, adding Tags can be async
	unlockOnce.Do(sd.RepositoryMutex.Unlock)

	for _, tag := range req.Tags {
		if _, err := sd.JobRepository.AddTagOrCreate(id, tag.Type, tag.Name); err != nil {
			log.Errorf("adding tag to new job %d failed: %s", id, err.Error())
			return
		}
	}

	log.Printf("new job (id: %d): cluster=%s, jobId=%d, user=%s, startTime=%d", id, req.Cluster, req.JobID, req.User, req.StartTime)
}

func (sd *SlurmNatsScheduler) checkAndHandleStopJob(job *schema.Job, req *StopJobRequest) {

	// Sanity checks
	if job == nil || job.StartTime.Unix() >= req.StopTime || job.State != schema.JobStateRunning {
		log.Errorf("stopTime must be larger than startTime and only running jobs can be stopped")
		return
	}

	if req.State != "" && !req.State.Valid() {
		log.Errorf("invalid job state: %#v", req.State)
		return
	} else if req.State == "" {
		req.State = schema.JobStateCompleted
	}

	// Mark job as stopped in the database (update state and duration)
	job.Duration = int32(req.StopTime - job.StartTime.Unix())
	job.State = req.State
	if err := sd.JobRepository.Stop(job.ID, job.Duration, job.State, job.MonitoringStatus); err != nil {
		log.Errorf("marking job as stopped failed: %s", err.Error())
		return
	}

	log.Printf("archiving job... (dbid: %d): cluster=%s, jobId=%d, user=%s, startTime=%s", job.ID, job.Cluster, job.JobID, job.User, job.StartTime)

	// Monitoring is disabled...
	if job.MonitoringStatus == schema.MonitoringStatusDisabled {
		return
	}

	// Trigger async archiving
	sd.JobRepository.TriggerArchiving(job)
}

func (sd *SlurmNatsScheduler) stopJob(req *StopJobRequest) {
	// if user := auth.GetUser(r.Context()); user != nil && !user.HasRole(auth.RoleApi) {
	// 	log.Errorf("missing role: %v", auth.GetRoleString(auth.RoleApi))
	// 	return
	// }

	log.Printf("Server Name: %s - Job ID: %v", *req.Cluster, req.JobId)

	// Fetch job (that will be stopped) from db
	var job *schema.Job
	var err error
	if req.JobId == nil {
		log.Errorf("the field 'jobId' is required")
		return
	}

	job, err = sd.JobRepository.Find(req.JobId, req.Cluster, req.StartTime)

	if err != nil {
		log.Errorf("finding job failed: %s", err.Error())
		return
	}

	sd.checkAndHandleStopJob(job, req)
}

func (sd *SlurmNatsScheduler) Init(rawConfig json.RawMessage) error {
	servers := []string{"nats://127.0.0.1:4222", "nats://127.0.0.1:1223"}

	nc, err := nats.Connect(strings.Join(servers, ","))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer nc.Close()

	getStatusTxt := func(nc *nats.Conn) string {
		switch nc.Status() {
		case nats.CONNECTED:
			return "Connected"
		case nats.CLOSED:
			return "Closed"
		default:
			return "Other"
		}
	}
	log.Printf("The connection status is %v\n", getStatusTxt(nc))

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer ec.Close()

	// Define the object
	type encodedMessage struct {
		ServerName   string
		ResponseCode int
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe
	if _, err := ec.Subscribe("test", func(s *encodedMessage) {
		log.Printf("Server Name: %s - Response Code: %v", s.ServerName, s.ResponseCode)
		if s.ResponseCode == 500 {
			wg.Done()
		}
	}); err != nil {
		log.Fatal(err.Error())
	}

	if _, err := ec.Subscribe("startJob", sd.startJob); err != nil {
		log.Fatal(err.Error())
	}

	if _, err := ec.Subscribe("stopJob", sd.stopJob); err != nil {
		log.Fatal(err.Error())
	}

	// Wait for a message to come in
	wg.Wait()

	return nil
}

func (sd *SlurmNatsScheduler) Sync() {

}
