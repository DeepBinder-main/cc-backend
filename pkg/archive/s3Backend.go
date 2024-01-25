// Copyright (C) 2023 NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package archive

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/ClusterCockpit/cc-backend/pkg/log"
	"github.com/ClusterCockpit/cc-backend/pkg/schema"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type S3ArchiveConfig struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"accessKeyID"`
	SecretAccessKey string `json:"secretAccessKey"`
	Bucket          string `json:"bucket"`
	UseSSL          bool   `json:"useSSL"`
}

type S3Archive struct {
	client   *minio.Client
	bucket   string
	clusters []string
}

func (s3a *S3Archive) Init(rawConfig json.RawMessage) (uint64, error) {
	var config S3ArchiveConfig
	var err error
	if err = json.Unmarshal(rawConfig, &config); err != nil {
		log.Warnf("Init() > Unmarshal error: %#v", err)
		return 0, err
	}

	fmt.Printf("Endpoint: %s Bucket: %s\n", config.Endpoint, config.Bucket)

	s3a.client, err = minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		err = fmt.Errorf("Init() : Initialize minio client failed")
		return 0, err
	}

	s3a.bucket = config.Bucket

	found, err := s3a.client.BucketExists(context.Background(), s3a.bucket)
	if err != nil {
		err = fmt.Errorf("Init() : %v", err)
		return 0, err
	}

	if found {
		log.Infof("Bucket found.")
	} else {
		log.Infof("Bucket not found.")
	}

	r, err := s3a.client.GetObject(context.Background(),
		s3a.bucket, "version.txt", minio.GetObjectOptions{})
	if err != nil {
		err = fmt.Errorf("Init() : Get version object failed")
		return 0, err
	}
	defer r.Close()

	b, err := io.ReadAll(r)
	if err != nil {
		log.Errorf("Init() : %v", err)
		return 0, err
	}
	version, err := strconv.ParseUint(strings.TrimSuffix(string(b), "\n"), 10, 64)
	if err != nil {
		log.Errorf("Init() : %v", err)
		return 0, err
	}

	if version != Version {
		return 0, fmt.Errorf("unsupported version %d, need %d", version, Version)
	}

	for object := range s3a.client.ListObjects(
		context.Background(),
		s3a.bucket, minio.ListObjectsOptions{
			Recursive: false,
		}) {

		if object.Err != nil {
			log.Errorf("listObject: %v", object.Err)
		}
		if strings.HasSuffix(object.Key, "/") {
			s3a.clusters = append(s3a.clusters, strings.TrimSuffix(object.Key, "/"))
		}
	}

	return version, err
}

func (s3a *S3Archive) Info() {
	fmt.Printf("Job archive %s\n", s3a.bucket)
	var clusters []string

	for object := range s3a.client.ListObjects(
		context.Background(),
		s3a.bucket, minio.ListObjectsOptions{
			Recursive: false,
		}) {

		if object.Err != nil {
			log.Errorf("listObject: %v", object.Err)
		}
		if strings.HasSuffix(object.Key, "/") {
			clusters = append(clusters, object.Key)
		}
	}
	ci := make(map[string]*clusterInfo)
	for _, cluster := range clusters {
		ci[cluster] = &clusterInfo{dateFirst: time.Now().Unix()}
		for d := range s3a.client.ListObjects(
			context.Background(),
			s3a.bucket, minio.ListObjectsOptions{
				Recursive: true,
				Prefix:    cluster,
			}) {
			log.Errorf("%s", d.Key)
			ci[cluster].diskSize += (float64(d.Size) * 1e-6)
		}
	}
}

//	func (s3a *S3Archive) Exists(job *schema.Job) bool {
//		return true
//	}

func (s3a *S3Archive) LoadJobMeta(job *schema.Job) (*schema.JobMeta, error) {
	filename := getPath(job, "/", "meta.json")
	log.Infof("Init() : %s", filename)

	r, err := s3a.client.GetObject(context.Background(),
		s3a.bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		err = fmt.Errorf("Init() : Get version object failed")
		return nil, err
	}
	b, err := io.ReadAll(r)
	if err != nil {
		log.Errorf("Init() : %v", err)
		return nil, err
	}

	return loadJobMeta(b)
}

func (s3a *S3Archive) LoadJobData(job *schema.Job) (schema.JobData, error) {
	var err error
	return schema.JobData{}, err
}

//	func (s3a *S3Archive) LoadClusterCfg(name string) (*schema.Cluster, error) {
//		var err error
//		return &schema.Cluster{}, err
//	}
//
// func (s3a *S3Archive) StoreJobMeta(jobMeta *schema.JobMeta) error
func (s3a *S3Archive) ImportJob(jobMeta *schema.JobMeta, jobData *schema.JobData) error {
	var err error
	return err
}

//
// func (s3a *S3Archive) GetClusters() []string
//
// func (s3a *S3Archive) CleanUp(jobs []*schema.Job)
//
// func (s3a *S3Archive) Move(jobs []*schema.Job, path string)
//
// func (s3a *S3Archive) Clean(before int64, after int64)
//
// func (s3a *S3Archive) Compress(jobs []*schema.Job)
//
// func (s3a *S3Archive) CompressLast(starttime int64) int64
//
// func (s3a *S3Archive) Iter(loadMetricData bool) <-chan JobContainer
