// Copyright (C) 2022 NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package archive

import (
	"encoding/json"

	"github.com/ClusterCockpit/cc-backend/pkg/log"
	"github.com/ClusterCockpit/cc-backend/pkg/schema"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3ArchiveConfig struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"accessKeyID"`
	SecretAccessKey string `json:"secretAccessKey"`
	UseSSL          bool   `json:"useSSL"`
}

type S3Archive struct {
	path string
}

func (s3a *S3Archive) Init(rawConfig json.RawMessage) (uint64, error) {
	var config S3ArchiveConfig
	if err := json.Unmarshal(rawConfig, &config); err != nil {
		log.Warnf("Init() > Unmarshal error: %#v", err)
		return 0, err
	}

	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return 0, err
}

func (s3a *S3Archive) Info() {
}

func (s3a *S3Archive) Exists(job *schema.Job) bool {
}

func (s3a *S3Archive) LoadJobMeta(job *schema.Job) (*schema.JobMeta, error) {
}

func (s3a *S3Archive) LoadJobData(job *schema.Job) (schema.JobData, error) {
}

func (s3a *S3Archive) LoadClusterCfg(name string) (*schema.Cluster, error) {
}

func (s3a *S3Archive) StoreJobMeta(jobMeta *schema.JobMeta) error

func (s3a *S3Archive) ImportJob(jobMeta *schema.JobMeta, jobData *schema.JobData) error

func (s3a *S3Archive) GetClusters() []string

func (s3a *S3Archive) CleanUp(jobs []*schema.Job)

func (s3a *S3Archive) Move(jobs []*schema.Job, path string)

func (s3a *S3Archive) Clean(before int64, after int64)

func (s3a *S3Archive) Compress(jobs []*schema.Job)

func (s3a *S3Archive) CompressLast(starttime int64) int64

func (s3a *S3Archive) Iter(loadMetricData bool) <-chan JobContainer
