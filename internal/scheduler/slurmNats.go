// Copyright (C) 2023 NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package scheduler

import (
	"encoding/json"
	"log"
	"strings"
	"sync"

	"github.com/nats-io/nats.go"
)

type SlurmNatsConfig struct {
	URL string `json:"url"`
}

type SlurmNatsScheduler struct {
	url string
}

func (sd *SlurmNatsScheduler) Init(rawConfig json.RawMessage) error {
	servers := []string{"nats://127.0.0.1:4222", "nats://127.0.0.1:1223"}

	nc, err := nats.Connect(strings.Join(servers, ","))
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
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
	if _, err := ec.Subscribe("stopJob", func(s *encodedMessage) {
		log.Printf("Server Name: %s - Response Code: %v", s.ServerName, s.ResponseCode)
		if s.ResponseCode == 500 {
			wg.Done()
		}
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()

	return nil
}

func (sd *SlurmNatsScheduler) Sync() {

}
