// Copyright (C) 2022 NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package repository

import (
	"sync"

	"github.com/Deepbinder-main/cc-backend/pkg/lrucache"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var (
	jobRepoOnce     sync.Once
	jobRepoInstance *JobRepository
)

type JobRepository struct {
	DB     *sqlx.DB
	driver string

	stmtCache *sq.StmtCache
	cache     *lrucache.Cache
}

func GetJobRepository() *JobRepository {
	jobRepoOnce.Do(func() {
		db := GetConnection()

		jobRepoInstance = &JobRepository{
			DB:     db.DB,
			driver: db.Driver,

			stmtCache: sq.NewStmtCache(db.DB),
			cache:     lrucache.New(1024 * 1024),
		}
		// // start archiving worker
		// go jobRepoInstance.archivingWorker()
	})
	return jobRepoInstance
}
