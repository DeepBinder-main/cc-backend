// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: create.sql

package db

import (
	"context"
)

const createFileStashURL = `-- name: CreateFileStashURL :exec
INSERT INTO file_stash_url (url) VALUES ($1)
`

func (q *Queries) CreateFileStashURL(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createFileStashURL)
	return err
}

const createInfluxDBConfiguration = `-- name: CreateInfluxDBConfiguration :exec
INSERT INTO influxdb_configurations (type, database_name, host, port, user, password, organization, ssl_enabled, batch_size, retry_interval, retry_exponential_base, max_retries, max_retry_time, meta_as_tags)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
`

func (q *Queries) CreateInfluxDBConfiguration(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createInfluxDBConfiguration)
	return err
}

const createLVMConf = `-- name: CreateLVMConf :exec
INSERT INTO lvm_conf (machine_id , username, minAvailableSpaceGB, maxAvailableSpaceGB) VALUES ($1, $2, $3 , $4)
`

func (q *Queries) CreateLVMConf(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createLVMConf)
	return err
}

const createLVStorageIssuer = `-- name: CreateLVStorageIssuer :exec
INSERT INTO lv_storage_issuer (machine_id, hostname, username, minAvailableSpaceGB, maxAvailableSpaceGB) VALUES ($1, $2, $3, $4, $5)
`

func (q *Queries) CreateLVStorageIssuer(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createLVStorageIssuer)
	return err
}

const createLogicalVolume = `-- name: CreateLogicalVolume :exec
INSERT INTO logical_volumes (machine_id, lv_name, vg_name, lv_attr, lv_size) VALUES ($1, $2, $3, $4, $5)
`

func (q *Queries) CreateLogicalVolume(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createLogicalVolume)
	return err
}

const createMachine = `-- name: CreateMachine :exec
INSERT INTO machines (machine_id, hostname, os_version, ip_address) VALUES ($1, $2, $3, $4)
`

func (q *Queries) CreateMachine(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createMachine)
	return err
}

const createMachineConf = `-- name: CreateMachineConf :exec
INSERT INTO machine_conf (machine_id, hostname, username, passphrase, port_number, password, host_key, folder_path) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

func (q *Queries) CreateMachineConf(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createMachineConf)
	return err
}

const createNotification = `-- name: CreateNotification :exec
INSERT INTO notifications (message) VALUES ($1)
`

func (q *Queries) CreateNotification(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createNotification)
	return err
}

const createPhysicalVolume = `-- name: CreatePhysicalVolume :exec
INSERT INTO physical_volumes (machine_id, pv_name, vg_name, pv_fmt, pv_attr, pv_size, pv_free) VALUES ($1, $2, $3, $4, $5, $6, $7)
`

func (q *Queries) CreatePhysicalVolume(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createPhysicalVolume)
	return err
}

const createRabbitMQConfig = `-- name: CreateRabbitMQConfig :exec
INSERT INTO rabbit_mq_config (conn_url, username, password) VALUES ($1, $2, $3)
`

func (q *Queries) CreateRabbitMQConfig(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createRabbitMQConfig)
	return err
}

const createRealtimeLog = `-- name: CreateRealtimeLog :exec
INSERT INTO realtime_logs (machine_id , log_message) VALUES ($1, $2)
`

func (q *Queries) CreateRealtimeLog(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createRealtimeLog)
	return err
}

const createVolumeGroup = `-- name: CreateVolumeGroup :exec
INSERT INTO volume_groups (machine_id, vg_name, pv_count, lv_count, snap_count, vg_attr, vg_size, vg_free) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

func (q *Queries) CreateVolumeGroup(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, createVolumeGroup)
	return err
}
