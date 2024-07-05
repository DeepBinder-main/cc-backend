-- Notifications
-- name: CreateNotification :exec
INSERT INTO notifications (message) VALUES (?);

-- name: GetNotifications :many
SELECT * FROM notifications
ORDER BY created_at DESC
LIMIT ?;

-- name: DeleteNotification :exec
DELETE FROM notifications WHERE id = ?;

-- Realtime Logs
-- name: CreateRealtimeLog :exec
INSERT INTO realtime_logs (log_message, machine_id) VALUES (?, ?);

-- name: GetRealtimeLogs :many
SELECT * FROM realtime_logs
WHERE machine_id = ?
ORDER BY created_at DESC
LIMIT ?;

-- name: DeleteRealtimeLog :exec
DELETE FROM realtime_logs WHERE id = ?;

-- LVM Conf
-- name: CreateLVMConf :exec
INSERT INTO lvm_conf (machine_id, username, minAvailableSpaceGB, maxAvailableSpaceGB)
VALUES (?, ?, ?, ?);

-- name: GetLVMConf :one
SELECT * FROM lvm_conf
WHERE machine_id = ?
ORDER BY created_at DESC
LIMIT 1;

-- name: UpdateLVMConf :exec
UPDATE lvm_conf
SET username = ?, minAvailableSpaceGB = ?, maxAvailableSpaceGB = ?
WHERE id = ?;

-- name: DeleteLVMConf :exec
DELETE FROM lvm_conf WHERE id = ?;

-- Machines
-- name: CreateMachine :exec
INSERT INTO machines (machine_id, hostname, os_version, ip_address)
VALUES (?, ?, ?, ?);

-- name: GetMachine :one
SELECT * FROM machines
WHERE machine_id = ?;

-- name: UpdateMachine :exec
UPDATE machines
SET hostname = ?, os_version = ?, ip_address = ?
WHERE machine_id = ?;

-- name: DeleteMachine :exec
DELETE FROM machines
WHERE machine_id = ?;

-- Logical Volumes
-- name: CreateLogicalVolume :exec
INSERT INTO logical_volumes (machine_id, lv_name, vg_name, lv_attr, lv_size)
VALUES (?, ?, ?, ?, ?);

-- name: GetLogicalVolumes :many
SELECT * FROM logical_volumes
WHERE machine_id = ?;

-- name: UpdateLogicalVolume :exec
UPDATE logical_volumes
SET lv_name = ?, vg_name = ?, lv_attr = ?, lv_size = ?
WHERE lv_id = ?;

-- name: DeleteLogicalVolume :exec
DELETE FROM logical_volumes
WHERE lv_id = ?;

-- Volume Groups
-- name: CreateVolumeGroup :exec
INSERT INTO volume_groups (machine_id, vg_name, pv_count, lv_count, snap_count, vg_attr, vg_size, vg_free)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetVolumeGroups :many
SELECT * FROM volume_groups
WHERE machine_id = ?;

-- name: UpdateVolumeGroup :exec
UPDATE volume_groups
SET vg_name = ?, pv_count = ?, lv_count = ?, snap_count = ?, vg_attr = ?, vg_size = ?, vg_free = ?
WHERE vg_id = ?;

-- name: DeleteVolumeGroup :exec
DELETE FROM volume_groups
WHERE vg_id = ?;

-- Physical Volumes
-- name: CreatePhysicalVolume :exec
INSERT INTO physical_volumes (machine_id, pv_name, vg_name, pv_fmt, pv_attr, pv_size, pv_free)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetPhysicalVolumes :many
SELECT * FROM physical_volumes
WHERE machine_id = ?;

-- name: UpdatePhysicalVolume :exec
UPDATE physical_volumes
SET pv_name = ?, vg_name = ?, pv_fmt = ?, pv_attr = ?, pv_size = ?, pv_free = ?
WHERE pv_id = ?;

-- name: DeletePhysicalVolume :exec
DELETE FROM physical_volumes
WHERE pv_id = ?;

-- LV Storage Issuer
-- name: CreateLVStorageIssuer :exec
INSERT INTO lv_storage_issuer (machine_id, inc_buffer, dec_buffer, hostname, username, minAvailableSpaceGB, maxAvailableSpaceGB)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetLVStorageIssuers :many
SELECT * FROM lv_storage_issuer;

-- name: UpdateLVStorageIssuer :exec
UPDATE lv_storage_issuer
SET inc_buffer = ?, dec_buffer = ?, hostname = ?, username = ?, minAvailableSpaceGB = ?, maxAvailableSpaceGB = ?
WHERE id = ?;

-- name: DeleteLVStorageIssuer :exec
DELETE FROM lv_storage_issuer WHERE id = ?;

-- Machine Conf
-- name: CreateMachineConf :exec
INSERT INTO machine_conf (machine_id, hostname, username, passphrase, port_number, password, host_key, folder_path)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetMachineConf :one
SELECT * FROM machine_conf
WHERE machine_id = ?;

-- name: UpdateMachineConf :exec
UPDATE machine_conf
SET hostname = ?, username = ?, passphrase = ?, port_number = ?, password = ?, host_key = ?, folder_path = ?
WHERE id = ?;

-- name: DeleteMachineConf :exec
DELETE FROM machine_conf WHERE id = ?;

-- File Stash URL
-- name: CreateFileStashURL :exec
INSERT INTO file_stash_url (url)
VALUES (?)
ON DUPLICATE KEY UPDATE url = VALUES(url);

-- name: GetFileStashURL :one
SELECT * FROM file_stash_url
LIMIT 1;

-- name: UpdateFileStashURL :exec
UPDATE file_stash_url
SET url = ?
WHERE single_row_enforcer = 1;

-- name: DeleteFileStashURL :exec
DELETE FROM file_stash_url WHERE id = ?;

-- RabbitMQ Config
-- name: CreateRabbitMQConfig :exec
INSERT INTO rabbit_mq_config (conn_url, username, password)
VALUES (?, ?, ?)
ON DUPLICATE KEY UPDATE conn_url = VALUES(conn_url), username = VALUES(username), password = VALUES(password);

-- name: GetRabbitMQConfig :one
SELECT * FROM rabbit_mq_config
LIMIT 1;

-- name: UpdateRabbitMQConfig :exec
UPDATE rabbit_mq_config
SET conn_url = ?, username = ?, password = ?
WHERE single_row_enforcer = 1;

-- name: DeleteRabbitMQConfig :exec
DELETE FROM rabbit_mq_config WHERE single_row_enforcer = 1;

-- InfluxDB Configurations
-- name: CreateInfluxDBConfiguration :exec
INSERT INTO influxdb_configurations (
    type, database_name, host, port, user, password, organization,
    ssl_enabled, batch_size, retry_interval, retry_exponential_base,
    max_retries, max_retry_time, meta_as_tags
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE
    type = VALUES(type),
    database_name = VALUES(database_name),
    host = VALUES(host),
    port = VALUES(port),
    user = VALUES(user),
    password = VALUES(password),
    organization = VALUES(organization),
    ssl_enabled = VALUES(ssl_enabled),
    batch_size = VALUES(batch_size),
    retry_interval = VALUES(retry_interval),
    retry_exponential_base = VALUES(retry_exponential_base),
    max_retries = VALUES(max_retries),
    max_retry_time = VALUES(max_retry_time),
    meta_as_tags = VALUES(meta_as_tags);

-- name: GetInfluxDBConfiguration :one
SELECT * FROM influxdb_configurations
LIMIT 1;

-- name: UpdateInfluxDBConfiguration :exec
UPDATE influxdb_configurations
SET type = ?, database_name = ?, host = ?, port = ?, user = ?, password = ?,
    organization = ?, ssl_enabled = ?, batch_size = ?, retry_interval = ?,
    retry_exponential_base = ?, max_retries = ?, max_retry_time = ?, meta_as_tags = ?
WHERE single_row_enforcer = 1;

-- name: DeleteInfluxDBConfiguration :exec
DELETE FROM influxdb_configurations WHERE id = ?;