// Copyright (C) 2022 NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ClusterCockpit/cc-backend/internal/scheduler"
	"github.com/nats-io/nats.go"
)

func usage() {
	log.Printf("Usage: nats-pub [-s server] [-creds file] <subject> <msg>\n")
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

func setupPublisher() {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	var userCreds = flag.String("creds", "", "User Credentials File")
	var showHelp = flag.Bool("h", false, "Show help message")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		showUsageAndExit(0)
	}

	args := flag.Args()
	if len(args) != 2 {
		showUsageAndExit(1)
	}

	fmt.Printf("Hello Nats\n")

	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Sample Publisher")}

	// Use UserCredentials
	if *userCreds != "" {
		opts = append(opts, nats.UserCredentials(*userCreds))
	}

	// Connect to NATS
	nc, err := nats.Connect(*urls, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	subj, msg := args[0], []byte(args[1])

	nc.Publish(subj, msg)
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}

	os.Exit(0)
}

func injectPayload() {
	// Read the JSON file
	jobsData, err := ioutil.ReadFile("slurm_0038.json")
	dbData, err := ioutil.ReadFile("slurmdb_0038-large.json")

	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Create an HTTP handler function
	http.HandleFunc("/slurm/v0.0.38/jobs", func(w http.ResponseWriter, r *http.Request) {
		// Set the response content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// Write the raw JSON data to the response writer
		_, err := w.Write(jobsData)
		if err != nil {
			http.Error(w, "Error writing jobsData payload", http.StatusInternalServerError)
			return
		}
	})

	http.HandleFunc("/slurmdb/v0.0.38/jobs", func(w http.ResponseWriter, r *http.Request) {
		// Set the response content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// Write the raw JSON data to the response writer
		_, err := w.Write(dbData)
		if err != nil {
			http.Error(w, "Error writing dbData payload", http.StatusInternalServerError)
			return
		}
	})

	// Start the HTTP server on port 8080
	fmt.Println("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}

func main() {
	cfgData := []byte(`{"target": "localhost"}`)

	var sch scheduler.SlurmNatsScheduler
	// sch.URL = "nats://127.0.0.1:1223"
	sch.Init(cfgData)

	// go injectPayload()

	os.Exit(0)
}
