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
	"time"

	"github.com/ClusterCockpit/cc-backend/pkg/log"
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

	apiEndpoint := "/slurmdb/v0.0.39/jobs"

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

	apiEndpoint := "/slurm/v0.0.39/jobs"
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
		time.Unix(int64(job["submit_time"].(float64)), 0).Format(time.RFC3339),
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

	latestJobs, err := queryAllJobs()
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, job := range latestJobs {
		// check if each job in latestJobs has existed combination of (job_id, cluster_id, start_time) in JobRepository
		jobs, err := sd.JobRepository.FindAll(&job.JobID, &job.Cluster, nil)
		if err != nil && err != sql.ErrNoRows {
			log.Errorf("checking for duplicate failed: %s", err.Error())
		} else if err == nil {
			// should update the JobRepository at this point
		}
	}

}
