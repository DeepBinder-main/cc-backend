// Copyright (C) 2023 NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package scheduler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/ClusterCockpit/cc-backend/pkg/log"
	"github.com/ClusterCockpit/cc-backend/pkg/schema"
)

// A Response struct to map the Entire Response
type Response struct {
	Name string `json:"name"`
	Jobs []Job  `json:"job_entries"`
}

type SlurmRestSchedulerConfig struct {
	URL string `json:"url"`
}

type SlurmRestScheduler struct {
	url string
}

var client *http.Client

func queryDB(qtime int64, clusterName string) ([]interface{}, error) {

	apiEndpoint := "/slurmdb/v0.0.38/jobs"

	// Construct the query parameters
	queryParams := url.Values{}
	queryParams.Set("users", "user1,user2")
	queryParams.Set("submit_time", "2023-01-01T00:00:00")

	// Add the query parameters to the API endpoint
	apiEndpoint += "?" + queryParams.Encode()

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		log.Errorf("Error creating request:", err)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Errorf("API request failed with status:", resp.Status)
	}

	// Read the response body
	// Here you can parse the response body as needed
	// For simplicity, let's just print the response body
	var dbOutput []byte
	_, err = resp.Body.Read(dbOutput)
	if err != nil {
		log.Errorf("Error reading response body:", err)
	}

	log.Errorf("API response:", string(dbOutput))

	dataJobs := make(map[string]interface{})
	err = json.Unmarshal(dbOutput, &dataJobs)
	if err != nil {
		log.Errorf("Error parsing JSON response:", err)
		os.Exit(1)
	}

	if _, ok := dataJobs["jobs"]; !ok {
		log.Errorf("ERROR: jobs not found - response incomplete")
		os.Exit(1)
	}

	jobs, _ := dataJobs["jobs"].([]interface{})
	return jobs, nil
}

func queryAllJobs() ([]interface{}, error) {
	var ctlOutput []byte

	apiEndpoint := "/slurm/v0.0.38/jobs"
	// Create a new HTTP GET request with query parameters
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		log.Errorf("Error creating request:", err)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Errorf("API request failed with status:", resp.Status)
	}

	_, err = resp.Body.Read(ctlOutput)
	if err != nil {
		log.Errorf("Error reading response body:", err)
	}

	dataJob := make(map[string]interface{})
	err = json.Unmarshal(ctlOutput, &dataJob)
	if err != nil {
		log.Errorf("Error parsing JSON response:", err)
		os.Exit(1)
	}

	if _, ok := dataJob["jobs"]; !ok {
		log.Errorf("ERROR: jobs not found - response incomplete")
		os.Exit(1)
	}

	jobs, _ := dataJob["jobs"].([]interface{})
	return jobs, nil
}

func printSlurmInfo(job map[string]interface{}) string {
	cpusPerTask := "1"
	tasksStr, ok := job["tasks"].(string)
	if !ok {
		tasksInt, _ := job["tasks"].(int)
		tasksStr = strconv.Itoa(tasksInt)
	}

	cpusStr, ok := job["cpus"].(string)
	if !ok {
		cpusInt, _ := job["cpus"].(int)
		cpusStr = strconv.Itoa(cpusInt)
	}

	tasks, _ := strconv.Atoi(tasksStr)
	cpus, _ := strconv.Atoi(cpusStr)
	if tasks > 0 {
		cpusPerTask = strconv.Itoa(int(math.Round(float64(cpus) / float64(tasks))))
	}

	text := fmt.Sprintf(`
	    JobId=%v JobName=%v
		UserId=%v(%v) GroupId=%v
		Account=%v QOS=%v
		Requeue=%v Restarts=%v BatchFlag=%v
		TimeLimit=%v
		SubmitTime=%v
		Partition=%v
		NodeList=%v
		NumNodes=%v NumCPUs=%v NumTasks=%v CPUs/Task=%v
		NTasksPerNode:Socket:Core=%v:%v:%v
		TRES_req=%v
		TRES_alloc=%v
		Command=%v
		WorkDir=%v
		StdErr=%v
		StdOut=%v`,
		job["job_id"], job["name"],
		job["user_name"], job["user_id"], job["group_id"],
		job["account"], job["qos"],
		job["requeue"], job["restart_cnt"], job["batch_flag"],
		job["time_limit"],
		time.Unix(int64(job["submit_time"].(float64)), 0).Format(time.RFC3338),
		job["partition"],
		job["nodes"],
		job["node_count"], cpus, tasks, cpusPerTask,
		job["tasks_per_node"], job["tasks_per_socket"], job["tasks_per_core"],
		job["tres_req_str"],
		job["tres_alloc_str"],
		job["command"],
		job["current_working_directory"],
		job["standard_error"],
		job["standard_output"],
	)

	return text
}

func exitWithError(err error, output []byte) {
	if exitError, ok := err.(*exec.ExitError); ok {
		if exitError.ExitCode() == 28 {
			fmt.Fprintf(os.Stderr, "ERROR: API call failed with timeout; check slurmrestd.\nOutput:\n%s\n", output)
		} else {
			fmt.Fprintf(os.Stderr, "ERROR: API call failed with code %d;\nOutput:\n%s\n", exitError.ExitCode(), output)
		}
	} else {
		log.Errorf("ERROR:", err)
	}
	os.Exit(1)
}

func loadClusterConfig(filename string) (map[string]interface{}, error) {
	clusterConfigData := make(map[string]interface{})

	file, err := os.Open(filename)
	if err != nil {
		log.Errorf("Cluster config file not found. No cores/GPU ids available.")
		return clusterConfigData, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&clusterConfigData)
	if err != nil {
		log.Errorf("Error decoding cluster config file:", err)
	}

	return clusterConfigData, err
}

func (sd *SlurmRestScheduler) Init(rawConfig json.RawMessage) error {
	clusterConfigData, err := loadClusterConfig("cluster-fritz.json")

	for k, v := range clusterConfigData {
		switch c := v.(type) {
		case string:
			fmt.Printf("Item %q is a string, containing %q\n", k, c)
		case float64:
			fmt.Printf("Looks like item %q is a number, specifically %f\n", k, c)
		default:
			fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, c)
		}
	}

	// Create an HTTP client
	client = &http.Client{}

	return err
}

func (sd *SlurmRestScheduler) Sync() {
	// for _, job := range jobs.GetJobs() {
	//     fmt.Printf("Job %s - %s\n", job.GetJobId(), job.GetJobState())
	// }

	jobsResponse, err := queryAllJobs()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Fetch an example instance of V0037JobsResponse
	// jobsResponse := V0037JobsResponse{}

	// Iterate over the Jobs slice
	for _, job := range jobsResponse.Jobs {
		// Process each job
		fmt.Printf("Job ID: %s\n", job.JobID)
		fmt.Printf("Job Name: %s\n", *job.Name)
		fmt.Printf("Job State: %s\n", *job.JobState)
		fmt.Println("Job StartTime:", *job.StartTime)

		// is aquire lock to avoid race condition between API calls needed?

		// aquire lock to avoid race condition between API calls
		var unlockOnce sync.Once
		sd.RepositoryMutex.Lock()
		defer unlockOnce.Do(sd.RepositoryMutex.Unlock)

		// is "running" one of JSON state?
		if *job.JobState == "running" {

			// Check if combination of (job_id, cluster_id, start_time) already exists:
			jobs, err := sd.JobRepository.FindAll(job.JobID, &job.Cluster, job.StartTime)

			if err != nil || err != sql.ErrNoRows {
				log.Errorf("checking for duplicate failed: %s", err.Error())
				return
			} else if err == nil {
				if len(jobs) == 0 {
					var defaultJob schema.BaseJob = schema.BaseJob{
						JobID:            job.JobID,
						User:             job.User,
						Project:          job.Project,
						Cluster:          job.Cluster,
						SubCluster:       job.SubCluster,
						Partition:        job.Partition,
						ArrayJobId:       job.ArrayJobId,
						NumNodes:         job.NumNodes,
						NumHWThreads:     job.NumHWThreads,
						NumAcc:           job.NumAcc,
						Exclusive:        job.Exclusive,
						MonitoringStatus: job.MonitoringStatus,
						SMT:              job.SMT,
						State:            job.State,
						Duration:         job.Duration,
						Walltime:         job.Walltime,
						Tags:             job.Tags,
						RawResources:     job.RawResources,
						Resources:        job.Resources,
						RawMetaData:      job.RawMetaData,
						MetaData:         job.MetaData,
						ConcurrentJobs:   job.ConcurrentJobs,
					}
					req := &schema.JobMeta{
						BaseJob:    defaultJob,
						StartTime:  job.StartTime,
						Statistics: make(map[string]schema.JobStatistics),
					}
					// req := new(schema.JobMeta)
					id, err := sd.JobRepository.Start(req)
				} else {
					for _, job := range jobs {
						log.Errorf("a job with that jobId, cluster and startTime already exists: dbid: %d", job.ID)
					}
				}
			}
		} else {
			// Check if completed job with combination of (job_id, cluster_id, start_time) already exists:
			existingJob, err := sd.JobRepository.Find(job.JobID, &job.Cluster, job.StartTime)

			if err == nil {
				existingJob.BaseJob.Duration = job.EndTime - job.StartTime
				existingJob.BaseJob.State = job.State
				existingJob.BaseJob.Walltime = job.StartTime
				req := &StopJobRequest{
					Cluster:   job.Cluster,
					JobId:     job.JobId,
					State:     job.State,
					StartTime: existingJob.StartTime,
					StopTime:  job.StartTime,
				}
				// req := new(schema.JobMeta)
				id, err := sd.JobRepository.checkAndHandleStopJob(job, req)
			}

		}
	}

}
