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

	return nil
}

func (sd *SlurmRestScheduler) Sync() {

}
