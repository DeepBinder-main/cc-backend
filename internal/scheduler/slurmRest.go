// Copyright (C) 2023 NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package scheduler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ClusterCockpit/cc-backend/internal/repository"
	"github.com/ClusterCockpit/cc-backend/pkg/log"
	"github.com/ClusterCockpit/cc-backend/pkg/schema"
)

type SlurmRestSchedulerConfig struct {
	URL string `json:"url"`

	JobRepository *repository.JobRepository

	clusterConfig ClusterConfig

	client *http.Client

	nodeClusterMap map[string]string

	clusterPCIEAddressMap map[string][]string
}

func (cfg *SlurmRestSchedulerConfig) queryDB(qtime int64, clusterName string) ([]interface{}, error) {

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
		log.Errorf("Error creating request: %v", err)
	}

	// Send the request
	resp, err := cfg.client.Do(req)
	if err != nil {
		log.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Errorf("API request failed with status: %v", resp.Status)
	}

	// Read the response body
	// Here you can parse the response body as needed
	// For simplicity, let's just print the response body
	var dbOutput []byte
	_, err = resp.Body.Read(dbOutput)
	if err != nil {
		log.Errorf("Error reading response body: %v", err)
	}

	log.Errorf("API response: %v", string(dbOutput))

	dataJobs := make(map[string]interface{})
	err = json.Unmarshal(dbOutput, &dataJobs)
	if err != nil {
		log.Errorf("Error parsing JSON response: %v", err)
		os.Exit(1)
	}

	if _, ok := dataJobs["jobs"]; !ok {
		log.Errorf("ERROR: jobs not found - response incomplete")
		os.Exit(1)
	}

	jobs, _ := dataJobs["jobs"].([]interface{})
	return jobs, nil
}

func (cfg *SlurmRestSchedulerConfig) fetchJobs() (SlurmPayload, error) {
	var ctlOutput []byte

	apiEndpoint := "http://:8080/slurm/v0.0.38/jobs"
	// Create a new HTTP GET request with query parameters
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		log.Errorf("Error creating request: %v", err)
	}

	// Send the request
	resp, err := cfg.client.Do(req)
	if err != nil {
		log.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		log.Errorf("API request failed with status: %v", resp.Status)
	}

	_, err = resp.Body.Read(ctlOutput)
	log.Printf("Received JSON Data: %v", ctlOutput)
	if err != nil {
		log.Errorf("Error reading response body: %v", err)
	}

	var jobsResponse SlurmPayload
	err = json.Unmarshal(ctlOutput, &jobsResponse)
	if err != nil {
		log.Errorf("Error parsing JSON response: %v", err)
		return jobsResponse, err
	}

	return jobsResponse, nil
}

func fetchJobsLocal() (SlurmPayload, error) {
	// Read the Slurm Payload JSON file
	jobsData, err := os.ReadFile("slurm_0038.json")

	if err != nil {
		fmt.Println("Error reading Slurm Payload JSON file:", err)
	}

	var jobsResponse SlurmPayload
	err = json.Unmarshal(jobsData, &jobsResponse)
	if err != nil {
		log.Errorf("Error parsing Slurm Payload JSON response: %v", err)
		return jobsResponse, err
	}

	return jobsResponse, nil
}

func fetchDumpedJobsLocal() (SlurmDBPayload, error) {
	// Read the SlurmDB Payload JSON file
	jobsData, err := os.ReadFile("slurmdb_0038-large.json")

	if err != nil {
		fmt.Println("Error reading SlurmDB Payload JSON file:", err)
	}

	var jobsResponse SlurmDBPayload
	err = json.Unmarshal(jobsData, &jobsResponse)
	if err != nil {
		log.Errorf("Error parsing SlurmDB Payload JSON response: %v", err)
		return jobsResponse, err
	}

	return jobsResponse, nil
}

func printSlurmInfo(job Job) string {

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
		job.JobID, job.Name,
		job.UserName, job.UserID, job.GroupID,
		job.Account, job.QoS,
		job.Requeue, job.RestartCnt, job.BatchFlag,
		job.TimeLimit, job.SubmitTime,
		job.Partition,
		job.Nodes,
		job.NodeCount, job.CPUs, job.Tasks, job.CPUPerTask,
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
		log.Errorf("ERROR: %v", err)
	}
	os.Exit(1)
}

func (cfg *SlurmRestSchedulerConfig) Init() error {
	var err error

	cfg.clusterConfig, err = DecodeClusterConfig("cluster-alex.json")

	cfg.nodeClusterMap = make(map[string]string)
	cfg.clusterPCIEAddressMap = make(map[string][]string)

	for _, subCluster := range cfg.clusterConfig.SubClusters {

		cfg.ConstructNodeClusterMap(subCluster.Nodes, subCluster.Name)

		pcieAddresses := make([]string, 0, 32)

		for idx, accelerator := range subCluster.Topology.Accelerators {
			pcieAddresses[idx] = accelerator.ID
		}

		cfg.clusterPCIEAddressMap[subCluster.Name] = pcieAddresses
	}
	// Create an HTTP client
	cfg.client = &http.Client{}

	return err
}

func (cfg *SlurmRestSchedulerConfig) checkAndHandleStopJob(job *schema.Job, req *StopJobRequest) {

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
	if err := cfg.JobRepository.Stop(job.ID, job.Duration, job.State, job.MonitoringStatus); err != nil {
		log.Errorf("marking job as stopped failed: %s", err.Error())
		return
	}

	log.Printf("archiving job... (dbid: %d): cluster=%s, jobId=%d, user=%s, startTime=%s", job.ID, job.Cluster, job.JobID, job.User, job.StartTime)

	// Monitoring is disabled...
	if job.MonitoringStatus == schema.MonitoringStatusDisabled {
		return
	}

	// Trigger async archiving
	cfg.JobRepository.TriggerArchiving(job)
}

func (cfg *SlurmRestSchedulerConfig) ConstructNodeClusterMap(nodes string, cluster string) {

	if cfg.nodeClusterMap == nil {
		cfg.nodeClusterMap = make(map[string]string)
	}

	// Split the input by commas
	groups := strings.Split(nodes, ",")

	for _, group := range groups {
		// Use regular expressions to match numbers and ranges
		numberRangeRegex := regexp.MustCompile(`a\[(\d+)-(\d+)\]`)
		numberRegex := regexp.MustCompile(`a(\d+)`)

		if numberRangeRegex.MatchString(group) {
			// Extract nodes from ranges
			matches := numberRangeRegex.FindStringSubmatch(group)
			if len(matches) == 3 {
				start, _ := strconv.Atoi(matches[1])
				end, _ := strconv.Atoi(matches[2])
				for i := start; i <= end; i++ {
					cfg.nodeClusterMap[matches[0]+fmt.Sprintf("%04d", i)] = cluster
				}
			}
		} else if numberRegex.MatchString(group) {
			// Extract individual node
			matches := numberRegex.FindStringSubmatch(group)
			if len(matches) == 2 {
				cfg.nodeClusterMap[group] = cluster
			}
		}
	}
}

func extractElements(indicesStr string, addresses []string) ([]string, error) {
	// Split the input string by commas to get individual index ranges
	indexRanges := strings.Split(indicesStr, ",")

	var selected []string
	for _, indexRange := range indexRanges {
		// Split each index range by hyphen to separate start and end indices
		rangeParts := strings.Split(indexRange, "-")

		if len(rangeParts) == 1 {
			// If there's only one part, it's a single index
			index, err := strconv.Atoi(rangeParts[0])
			if err != nil {
				return nil, err
			}
			selected = append(selected, addresses[index])
		} else if len(rangeParts) == 2 {
			// If there are two parts, it's a range
			start, err := strconv.Atoi(rangeParts[0])
			if err != nil {
				return nil, err
			}
			end, err := strconv.Atoi(rangeParts[1])
			if err != nil {
				return nil, err
			}
			// Add all indices in the range to the result
			for i := start; i <= end; i++ {
				selected = append(selected, addresses[i])
			}
		} else {
			// Invalid format
			return nil, fmt.Errorf("invalid index range: %s", indexRange)
		}
	}

	return selected, nil
}

func (cfg *SlurmRestSchedulerConfig) CreateJobMeta(job Job) (*schema.JobMeta, error) {

	var exclusive int32
	if job.Shared == nil {
		exclusive = 1
	} else {
		exclusive = 0
	}

	totalGPUs := 0

	var resources []*schema.Resource

	for nodeIndex, node := range job.JobResources.AllocatedNodes {
		var res schema.Resource
		res.Hostname = node.Nodename

		log.Debugf("Node %s Cores map size: %d\n", node.Nodename, len(node.Sockets))

		if node.CPUsUsed == nil || node.MemoryAllocated == nil {
			log.Fatalf("Either node.Cpus or node.Memory is nil\n")
		}

		for k, v := range node.Sockets {
			fmt.Printf("core id[%s] value[%s]\n", k, v)
			threadID, _ := strconv.Atoi(k)
			res.HWThreads = append(res.HWThreads, threadID)
		}

		re := regexp.MustCompile(`\(([^)]*)\)`)
		matches := re.FindStringSubmatch(job.GresDetail[nodeIndex])

		if len(matches) < 2 {
			return nil, fmt.Errorf("no substring found in brackets")
		}

		nodePCIEAddresses := cfg.clusterPCIEAddressMap[cfg.nodeClusterMap[node.Nodename]]

		selectedPCIEAddresses, err := extractElements(matches[1], nodePCIEAddresses)

		totalGPUs += len(selectedPCIEAddresses)

		if err != nil {
			return nil, err
		}

		// For core/GPU id mapping, need to query from cluster config file
		res.Accelerators = selectedPCIEAddresses
		resources = append(resources, &res)
	}

	metaData := make(map[string]string)
	metaData["jobName"] = job.Name
	metaData["slurmInfo"] = printSlurmInfo(job)

	metaDataInBytes, err := json.Marshal(metaData)
	if err != nil {
		log.Fatalf("metaData JSON marshaling failed: %s", err)
	}

	var defaultJob schema.BaseJob = schema.BaseJob{
		JobID:     job.JobID,
		User:      job.UserName,
		Project:   job.Account,
		Cluster:   job.Cluster,
		Partition: job.Partition,
		// check nil
		ArrayJobId:   job.ArrayJobID,
		NumNodes:     job.NodeCount,
		NumHWThreads: job.CPUs,
		NumAcc:       int32(totalGPUs),
		Exclusive:    exclusive,
		// MonitoringStatus: job.MonitoringStatus,
		// SMT:            job.TasksPerCore,
		State: schema.JobState(job.JobState),
		// ignore this for start job
		// Duration:       int32(time.Now().Unix() - job.StartTime), // or SubmitTime?
		Walltime: time.Now().Unix(), // max duration requested by the job
		// Tags:           job.Tags,
		// ignore this!
		// RawResources: jobResourcesInBytes,
		// "job_resources": "allocated_nodes" "sockets":
		// very important; has to be right
		Resources:   resources,
		RawMetaData: metaDataInBytes,
		// optional metadata with'jobScript 'jobName': 'slurmInfo':
		MetaData: metaData,
		// ConcurrentJobs: job.ConcurrentJobs,
	}
	log.Debugf("Generated BaseJob with Resources=%v", defaultJob.Resources[0])

	meta := &schema.JobMeta{
		BaseJob:    defaultJob,
		StartTime:  job.StartTime,
		Statistics: make(map[string]schema.JobStatistics),
	}
	// log.Debugf("Generated JobMeta %v", req.BaseJob.JobID)

	return meta, nil
}

func (cfg *SlurmRestSchedulerConfig) HandleJobs(jobs []Job) error {

	// runningJobsInCC, err := cfg.JobRepository.FindRunningJobs("alex")

	// Iterate over the Jobs slice
	for _, job := range jobs {
		// Process each job from Slurm
		fmt.Printf("Job ID: %d\n", job.JobID)
		fmt.Printf("Job Name: %s\n", job.Name)
		fmt.Printf("Job State: %s\n", job.JobState)
		fmt.Println("Job StartTime:", job.StartTime)
		fmt.Println("Job Cluster:", job.Cluster)

		if job.JobState == "RUNNING" {

			meta, _ := cfg.CreateJobMeta(job)

			// For all running jobs from Slurm
			_, notFoundError := cfg.JobRepository.Find(&job.JobID, &job.Cluster, &job.StartTime)

			if notFoundError != nil {
				// if it does not exist in CC, create a new entry
				log.Print("Job does not exist in CC, will create a new entry:", job.JobID)
				id, startJobError := cfg.JobRepository.Start(meta)
				if startJobError != nil {
					return startJobError
				}
				log.Debug("Added job", id)
			}

			// Running in both sides: nothing needs to be done
		} else if job.JobState == "COMPLETED" {
			// Check if completed job with combination of (job_id, cluster_id, start_time) already exists
			log.Debugf("Processing completed job ID: %v Cluster: %v StartTime: %v", job.JobID, job.Cluster, job.StartTime)
			existingJob, err := cfg.JobRepository.Find(&job.JobID, &job.Cluster, &job.StartTime)

			if err == nil && existingJob.State != schema.JobStateCompleted {
				// for jobs completed in Slurm (either in Slurm or maybe SlurmDB)
				// update job in CC with new info (job final status, duration, end timestamp)

				existingJob.BaseJob.Duration = int32(job.EndTime - job.StartTime)
				existingJob.BaseJob.State = schema.JobState(job.JobState)
				existingJob.BaseJob.Walltime = job.EndTime

				req := &StopJobRequest{
					Cluster:   &job.Cluster,
					JobId:     &job.JobID,
					State:     schema.JobState(job.JobState),
					StartTime: &job.StartTime,
					StopTime:  job.EndTime,
				}
				cfg.checkAndHandleStopJob(existingJob, req)
			}
		}
	}
	return nil
}

func (cfg *SlurmRestSchedulerConfig) HandleDumpedJobs(jobs []DumpedJob) error {

	// Iterate over the Jobs slice
	for _, job := range jobs {
		// Process each job from Slurm
		fmt.Printf("Job ID: %d\n", job.JobID)
		fmt.Printf("Job Name: %s\n", job.Name)
		fmt.Printf("Job State: %s\n", job.State.Current)
		fmt.Println("Job EndTime:", job.Time.End)
		fmt.Println("Job Cluster:", job.Cluster)

		// Check if completed job with combination of (job_id, cluster_id, start_time) already exists
		log.Debugf("Processing completed dumped job ID: %v Cluster: %v StartTime: %v", job.JobID, job.Cluster, job.Time.Start)
		existingJob, err := cfg.JobRepository.Find(&job.JobID, &job.Cluster, &job.Time.Start)

		if err == nil && existingJob.State != schema.JobStateCompleted {
			// for jobs completed in Slurm (either in Slurm or maybe SlurmDB)
			// update job in CC with new info (job final status, duration, end timestamp)

			existingJob.BaseJob.Duration = int32(job.Time.End - job.Time.Start)
			existingJob.BaseJob.State = schema.JobState(job.State.Current)
			existingJob.BaseJob.Walltime = job.Time.End

			req := &StopJobRequest{
				Cluster:   &job.Cluster,
				JobId:     &job.JobID,
				State:     schema.JobState(job.State.Current),
				StartTime: &job.Time.Start,
				StopTime:  job.Time.End,
			}
			cfg.checkAndHandleStopJob(existingJob, req)
		}
	}
	return nil
}

func (cfg *SlurmRestSchedulerConfig) Sync() {

	// Fetch an instance of Slurm JobsResponse
	jobsResponse, err := fetchJobsLocal()
	if err != nil {
		log.Fatal(err.Error())
	}
	cfg.HandleJobs(jobsResponse.Jobs)

	// Fetch an instance of Slurm DB JobsResponse
	dumpedJobsResponse, err := fetchDumpedJobsLocal()
	if err != nil {
		log.Fatal(err.Error())
	}
	cfg.HandleDumpedJobs(dumpedJobsResponse.Jobs)

}
