// Copyright (C) 2023 NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package scheduler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	"github.com/ClusterCockpit/cc-backend/internal/repository"
	"github.com/ClusterCockpit/cc-backend/pkg/log"
	"github.com/ClusterCockpit/cc-backend/pkg/schema"

	openapi "github.com/ClusterCockpit/slurm-rest-client-0_0_38"
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

	JobRepository *repository.JobRepository
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

func printSlurmInfo(job openapi.V0038JobResponseProperties) string {

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
		job.JobId, job.Name,
		job.UserName, job.UserId, job.GroupId,
		job.Account, job.Qos,
		job.Requeue, job.RestartCnt, job.BatchFlag,
		job.TimeLimit, job.SubmitTime,
		//time.Unix(int64(*.(float64)), 0).Format(time.RFC1123),
		job.Partition,
		job.Nodes,
		job.NodeCount, job.Cpus, job.Tasks, job.CpusPerTask,
		job.TasksPerBoard, job.TasksPerSocket, job.TasksPerCore,
		job.TresAllocStr,
		job.TresAllocStr,
		job.Command,
		job.CurrentWorkingDirectory,
		job.StandardError,
		job.StandardOutput,
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

func (sd *SlurmRestScheduler) checkAndHandleStopJob(job *schema.Job, req *StopJobRequest) {

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

func (sd *SlurmRestScheduler) HandleJobsResponse(jobsResponse openapi.V0038JobsResponse) {

	// Iterate over the Jobs slice
	for _, job := range jobsResponse.Jobs {
		// Process each job
		fmt.Printf("Job ID: %s\n", job.JobId)
		fmt.Printf("Job Name: %s\n", *job.Name)
		fmt.Printf("Job State: %s\n", *job.JobState)
		fmt.Println("Job StartTime:", *job.StartTime)

		// aquire lock to avoid race condition between API calls
		// var unlockOnce sync.Once
		// sd.RepositoryMutex.Lock()
		// defer unlockOnce.Do(sd.RepositoryMutex.Unlock)

		// is "running" one of JSON state?
		if *job.JobState == "running" {

			jobs, err := sd.JobRepository.FindRunningJobs(*job.Cluster)
			if err != nil {
				log.Fatalf("Failed to find running jobs: %v", err)
			}

			for id, job := range jobs {
				fmt.Printf("Job ID: %d, Job: %+v\n", id, job)
			}

			if err != nil || err != sql.ErrNoRows {
				log.Errorf("checking for duplicate failed: %s", err.Error())
				return
			} else if err == nil {
				if len(jobs) == 0 {
					var exclusive int32
					if job.Shared == nil {
						exclusive = 1
					} else {
						exclusive = 0
					}

					jobResourcesInBytes, err := json.Marshal(*job.JobResources)
					if err != nil {
						log.Fatalf("JSON marshaling failed: %s", err)
					}

					var resources []*schema.Resource

					// Define a regular expression to match "gres/gpu=x"
					regex := regexp.MustCompile(`gres/gpu=(\d+)`)

					// Find all matches in the input string
					matches := regex.FindAllStringSubmatch(*job.TresAllocStr, -1)

					// Initialize a variable to store the total number of GPUs
					var totalGPUs int32
					// Iterate through the matches
					match := matches[0]
					if len(match) == 2 {
						gpuCount, _ := strconv.Atoi(match[1])
						totalGPUs += int32(gpuCount)
					}

					for _, node := range job.JobResources.AllocatedNodes {
						var res schema.Resource
						res.Hostname = *node.Nodename
						for k, v := range node.Sockets.Cores {
							fmt.Printf("core id[%s] value[%s]\n", k, v)
							threadID, _ := strconv.Atoi(k)
							res.HWThreads = append(res.HWThreads, threadID)
						}
						res.Accelerators = append(res.Accelerators, *job.TresAllocStr)
						// cpu=512,mem=1875G,node=4,billing=512,gres\/gpu=32,gres\/gpu:a40=32
						resources = append(resources, &res)
					}

					var metaData map[string]string
					metaData["jobName"] = *job.Name
					metaData["slurmInfo"] = printSlurmInfo(job)
					// metaData["jobScript"] = "What to put here?"
					metaDataInBytes, err := json.Marshal(metaData)

					var defaultJob schema.BaseJob = schema.BaseJob{
						JobID:     int64(*job.JobId),
						User:      *job.UserName,
						Project:   *job.Account,
						Cluster:   *job.Cluster,
						Partition: *job.Partition,
						// check nil
						ArrayJobId:   int64(*job.ArrayJobId),
						NumNodes:     *job.NodeCount,
						NumHWThreads: *job.Cpus,
						NumAcc:       totalGPUs,
						Exclusive:    exclusive,
						// MonitoringStatus: job.MonitoringStatus,
						// SMT:            *job.TasksPerCore,
						State: schema.JobState(*job.JobState),
						// ignore this for start job
						// Duration:       int32(time.Now().Unix() - *job.StartTime), // or SubmitTime?
						Walltime: time.Now().Unix(), // max duration requested by the job
						// Tags:           job.Tags,
						// ignore this!
						RawResources: jobResourcesInBytes,
						// "job_resources": "allocated_nodes" "sockets":
						// very important; has to be right
						Resources:   resources,
						RawMetaData: metaDataInBytes,
						// optional metadata with'jobScript 'jobName': 'slurmInfo':
						MetaData: metaData,
						// ConcurrentJobs: job.ConcurrentJobs,
					}
					req := &schema.JobMeta{
						BaseJob:    defaultJob,
						StartTime:  *job.StartTime,
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
			var jobID int64
			jobID = int64(*job.JobId)
			existingJob, err := sd.JobRepository.Find(&jobID, job.Cluster, job.StartTime)

			if err == nil {
				existingJob.BaseJob.Duration = int32(*job.EndTime - *job.StartTime)
				existingJob.BaseJob.State = schema.JobState(*job.JobState)
				existingJob.BaseJob.Walltime = *job.StartTime

				req := &StopJobRequest{
					Cluster:   job.Cluster,
					JobId:     &jobID,
					State:     schema.JobState(*job.JobState),
					StartTime: &existingJob.StartTimeUnix,
					StopTime:  *job.EndTime,
				}
				// req := new(schema.JobMeta)
				sd.checkAndHandleStopJob(existingJob, req)
			}

		}
	}
}

func (sd *SlurmRestScheduler) Sync() {
	// for _, job := range jobs.GetJobs() {
	//     fmt.Printf("Job %s - %s\n", job.GetJobId(), job.GetJobState())
	// }

	response, err := queryAllJobs()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Fetch an example instance of V0037JobsResponse
	// jobsResponse := openapi.V0038JobsResponse{}

	var jobsResponse openapi.V0038JobsResponse
	sd.HandleJobsResponse(jobsResponse)

}
