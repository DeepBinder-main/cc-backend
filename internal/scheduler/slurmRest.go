// Copyright (C) 2023 NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package scheduler

import "encoding/json"

type SlurmRestSchedulerConfig struct {
	URL string `json:"url"`
}

type SlurmRestScheduler struct {
	url string
}

func (sd *SlurmRestScheduler) Init(rawConfig json.RawMessage) error {
	// cfg := slurmrest.NewConfiguration()
	// cfg.HTTPClient = &http.Client{Timeout: time.Second * 3600}
	// cfg.Scheme = "http"
	// cfg.Host = "localhost"

	// client := slurmrest.NewAPIClient(cfg)
	return nil
}

func (sd *SlurmRestScheduler) Sync() {
	// jreq := client.SlurmApi.SlurmctldGetJobs(context.Background())
	// jobs, resp, err := client.SlurmApi.SlurmctldGetJobsExecute(jreq)
	// if err != nil {
	// 	log.Fatalf("FAIL: %s", err)
	// } else if resp.StatusCode != 200 {
	// 	log.Fatalf("Invalid status code: %d\n", resp.StatusCode)
	// }

	// for _, job := range jobs.GetJobs() {
	//     fmt.Printf("Job %s - %s\n", job.GetJobId(), job.GetJobState())
	// }
}
