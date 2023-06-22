// Copyright (C) 2023 NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package scheduler

import "encoding/json"

type BatchScheduler interface {
	Init(rawConfig json.RawMessage) error

	Sync()
}

var sd BatchScheduler

func Init(rawConfig json.RawMessage) error {

	sd = &SlurmNatsScheduler{}
	sd.Init(rawConfig)

	return nil
}

func GetHandle() BatchScheduler {
	return sd
}
