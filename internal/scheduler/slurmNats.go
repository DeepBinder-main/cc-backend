// Copyright (C) 2023 NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package scheduler

import "encoding/json"

type SlurmNatsConfig struct {
	URL string `json:"url"`
}

type SlurmNatsScheduler struct {
	url string
}

func (sd *SlurmNatsScheduler) Init(rawConfig json.RawMessage) error {

	return nil
}

func (sd *SlurmNatsScheduler) Sync() {

}
