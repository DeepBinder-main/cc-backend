#!/bin/bash

# Directories
BASE_DIR="internal/repository/sqlc/queries"
NOTIFICATIONS_DIR="${BASE_DIR}/notifications"
INFLUXDB_CONFIGURATIONS_DIR="${BASE_DIR}/influxdb_configurations"
REALTIME_LOGS_DIR="${BASE_DIR}/realtime_logs"
LVM_CONF_DIR="${BASE_DIR}/lvm_conf"
MACHINES_DIR="${BASE_DIR}/machines"
LOGICAL_VOLUMES_DIR="${BASE_DIR}/logical_volumes"
VOLUME_GROUPS_DIR="${BASE_DIR}/volume_groups"
PHYSICAL_VOLUMES_DIR="${BASE_DIR}/physical_volumes"
LV_STORAGE_ISSUER_DIR="${BASE_DIR}/lv_storage_issuer"
MACHINE_CONF_DIR="${BASE_DIR}/machine_conf"
FILE_STASH_URL_DIR="${BASE_DIR}/file_stash_url"
RABBIT_MQ_CONFIG_DIR="${BASE_DIR}/rabbit_mq_config"

# Create directories
mkdir -p "$NOTIFICATIONS_DIR"
mkdir -p "$INFLUXDB_CONFIGURATIONS_DIR"
mkdir -p "$REALTIME_LOGS_DIR"
mkdir -p "$LVM_CONF_DIR"
mkdir -p "$MACHINES_DIR"
mkdir -p "$LOGICAL_VOLUMES_DIR"
mkdir -p "$VOLUME_GROUPS_DIR"
mkdir -p "$PHYSICAL_VOLUMES_DIR"
mkdir -p "$LV_STORAGE_ISSUER_DIR"
mkdir -p "$MACHINE_CONF_DIR"
mkdir -p "$FILE_STASH_URL_DIR"
mkdir -p "$RABBIT_MQ_CONFIG_DIR"

# Create .sql files for each table
# Notifications
cat << EOF > "$NOTIFICATIONS_DIR/create.sql"
-- name: CreateNotification :exec
INSERT INTO notifications (message) VALUES (\$1);
EOF

cat << EOF > "$NOTIFICATIONS_DIR/read.sql"
-- name: GetNotificationByID :one
SELECT id, message, created_at FROM notifications WHERE id = \$1;

-- name: ListNotifications :many
SELECT id, message, created_at FROM notifications ORDER BY created_at DESC;
EOF

cat << EOF > "$NOTIFICATIONS_DIR/delete.sql"
-- name: DeleteNotificationByID :exec
DELETE FROM notifications WHERE id = \$1;
EOF

# InfluxDB Configurations
cat << EOF > "$INFLUXDB_CONFIGURATIONS_DIR/create.sql"
-- name: CreateInfluxDBConfiguration :exec
INSERT INTO influxdb_configurations (type, database_name, host, port, user, password, organization, ssl_enabled, batch_size, retry_interval, retry_exponential_base, max_retries, max_retry_time, meta_as_tags)
VALUES (\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8, \$9, \$10, \$11, \$12, \$13, \$14);
EOF

cat << EOF > "$INFLUXDB_CONFIGURATIONS_DIR/read.sql"
-- name: GetInfluxDBConfigurationByID :one
SELECT * FROM influxdb_configurations WHERE id = \$1;

-- name: ListInfluxDBConfigurations :many
SELECT * FROM influxdb_configurations ORDER BY id;
EOF

cat << EOF > "$INFLUXDB_CONFIGURATIONS_DIR/delete.sql"
-- name: DeleteInfluxDBConfigurationByID :exec
DELETE FROM influxdb_configurations WHERE id = \$1;
EOF

# Realtime Logs
cat << EOF > "$REALTIME_LOGS_DIR/create.sql"
-- name: CreateRealtimeLog :exec
INSERT INTO realtime_logs (machine_id , log_message) VALUES (\$1, \$2);
EOF

cat << EOF > "$REALTIME_LOGS_DIR/read.sql"
-- name: GetRealtimeLogByID :one
SELECT id,machine_id ,log_message, created_at FROM realtime_logs WHERE id = \$1;

-- name: ListRealtimeLogs :many
SELECT id,machine_id ,log_message, created_at FROM realtime_logs ORDER BY created_at DESC;
EOF

cat << EOF > "$REALTIME_LOGS_DIR/delete.sql"
-- name: DeleteRealtimeLogByID :exec
DELETE FROM realtime_logs WHERE id = \$1;
EOF

# LVM Conf
cat << EOF > "$LVM_CONF_DIR/create.sql"
-- name: CreateLVMConf :exec
INSERT INTO lvm_conf (machine_id , username, minAvailableSpaceGB, maxAvailableSpaceGB) VALUES (\$1, \$2, \$3 , \$4);
EOF

cat << EOF > "$LVM_CONF_DIR/read.sql"
-- name: GetLVMConfByID :one
SELECT id,machine_id , username, minAvailableSpaceGB, maxAvailableSpaceGB FROM lvm_conf WHERE id = \$1;

-- name: ListLVMConfs :many
SELECT id, machine_id, username, minAvailableSpaceGB, maxAvailableSpaceGB FROM lvm_conf ORDER BY id;
EOF

cat << EOF > "$LVM_CONF_DIR/delete.sql"
-- name: DeleteLVMConfByID :exec
DELETE FROM lvm_conf WHERE id = \$1;
EOF

# Machines
cat << EOF > "$MACHINES_DIR/create.sql"
-- name: CreateMachine :exec
INSERT INTO machines (machine_id, hostname, os_version, ip_address) VALUES (\$1, \$2, \$3, \$4);
EOF

cat << EOF > "$MACHINES_DIR/read.sql"
-- name: GetMachineByID :one
SELECT machine_id, hostname, os_version, ip_address FROM machines WHERE machine_id = \$1;

-- name: ListMachines :many
SELECT machine_id, hostname, os_version, ip_address FROM machines ORDER BY machine_id;
EOF

cat << EOF > "$MACHINES_DIR/delete.sql"
-- name: DeleteMachineByID :exec
DELETE FROM machines WHERE machine_id = \$1;
EOF

# Logical Volumes
cat << EOF > "$LOGICAL_VOLUMES_DIR/create.sql"
-- name: CreateLogicalVolume :exec
INSERT INTO logical_volumes (machine_id, lv_name, vg_name, lv_attr, lv_size) VALUES (\$1, \$2, \$3, \$4, \$5);
EOF

cat << EOF > "$LOGICAL_VOLUMES_DIR/read.sql"
-- name: GetLogicalVolumeByID :one
SELECT lv_id, machine_id, lv_name, vg_name, lv_attr, lv_size FROM logical_volumes WHERE lv_id = \$1;

-- name: ListLogicalVolumes :many
SELECT lv_id, machine_id, lv_name, vg_name, lv_attr, lv_size FROM logical_volumes ORDER BY lv_id;
EOF

cat << EOF > "$LOGICAL_VOLUMES_DIR/delete.sql"
-- name: DeleteLogicalVolumeByID :exec
DELETE FROM logical_volumes WHERE lv_id = \$1;
EOF

# Volume Groups
cat << EOF > "$VOLUME_GROUPS_DIR/create.sql"
-- name: CreateVolumeGroup :exec
INSERT INTO volume_groups (machine_id, vg_name, pv_count, lv_count, snap_count, vg_attr, vg_size, vg_free) VALUES (\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8);
EOF

cat << EOF > "$VOLUME_GROUPS_DIR/read.sql"
-- name: GetVolumeGroupByID :one
SELECT vg_id, machine_id, vg_name, pv_count, lv_count, snap_count, vg_attr, vg_size, vg_free FROM volume_groups WHERE vg_id = \$1;

-- name: ListVolumeGroups :many
SELECT vg_id, machine_id, vg_name, pv_count, lv_count, snap_count, vg_attr, vg_size, vg_free FROM volume_groups ORDER BY vg_id;
EOF

cat << EOF > "$VOLUME_GROUPS_DIR/delete.sql"
-- name: DeleteVolumeGroupByID :exec
DELETE FROM volume_groups WHERE vg_id = \$1;
EOF

# Physical Volumes
cat << EOF > "$PHYSICAL_VOLUMES_DIR/create.sql"
-- name: CreatePhysicalVolume :exec
INSERT INTO physical_volumes (machine_id, pv_name, vg_name, pv_fmt, pv_attr, pv_size, pv_free) VALUES (\$1, \$2, \$3, \$4, \$5, \$6, \$7);
EOF

cat << EOF > "$PHYSICAL_VOLUMES_DIR/read.sql"
-- name: GetPhysicalVolumeByID :one
SELECT pv_id, machine_id, pv_name, vg_name, pv_fmt, pv_attr, pv_size, pv_free FROM physical_volumes WHERE pv_id = \$1;

-- name: ListPhysicalVolumes :many
SELECT pv_id, machine_id, pv_name, vg_name, pv_fmt, pv_attr, pv_size, pv_free FROM physical_volumes ORDER BY pv_id;
EOF

cat << EOF > "$PHYSICAL_VOLUMES_DIR/delete.sql"
-- name: DeletePhysicalVolumeByID :exec
DELETE FROM physical_volumes WHERE pv_id = \$1;
EOF

# LV Storage Issuer
cat << EOF > "$LV_STORAGE_ISSUER_DIR/create.sql"
-- name: CreateLVStorageIssuer :exec
INSERT INTO lv_storage_issuer (machine_id, hostname, username, minAvailableSpaceGB, maxAvailableSpaceGB) VALUES (\$1, \$2, \$3, \$4, \$5);
EOF

cat << EOF > "$LV_STORAGE_ISSUER_DIR/read.sql"
-- name: GetLVStorageIssuerByID :one
SELECT id, machine_id, hostname, username, minAvailableSpaceGB, maxAvailableSpaceGB FROM lv_storage_issuer WHERE id = \$1;

-- name: ListLVStorageIssuers :many
SELECT id, machine_id, hostname, username, minAvailableSpaceGB, maxAvailableSpaceGB FROM lv_storage_issuer ORDER BY id;
EOF

cat << EOF > "$LV_STORAGE_ISSUER_DIR/delete.sql"
-- name: DeleteLVStorageIssuerByID :exec
DELETE FROM lv_storage_issuer WHERE id = \$1;
EOF

# Machine Conf
cat << EOF > "$MACHINE_CONF_DIR/create.sql"
-- name: CreateMachineConf :exec
INSERT INTO machine_conf (machine_id, hostname, username, passphrase, port_number, password, host_key, folder_path) VALUES (\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8);
EOF

cat << EOF > "$MACHINE_CONF_DIR/read.sql"
-- name: GetMachineConfByID :one
SELECT id, machine_id, hostname, username, passphrase, port_number, password, host_key, folder_path FROM machine_conf WHERE id = \$1;

-- name: ListMachineConfs :many
SELECT id, machine_id, hostname, username, passphrase, port_number, password, host_key, folder_path FROM machine_conf ORDER BY id;
EOF

cat << EOF > "$MACHINE_CONF_DIR/delete.sql"
-- name: DeleteMachineConfByID :exec
DELETE FROM machine_conf WHERE id = \$1;
EOF

# File Stash URL
cat << EOF > "$FILE_STASH_URL_DIR/create.sql"
-- name: CreateFileStashURL :exec
INSERT INTO file_stash_url (url) VALUES (\$1);
EOF

cat << EOF > "$FILE_STASH_URL_DIR/read.sql"
-- name: GetFileStashURLByID :one
SELECT id, url, created_at FROM file_stash_url WHERE id = \$1;

-- name: ListFileStashURLs :many
SELECT id, url, created_at FROM file_stash_url ORDER BY created_at DESC;
EOF

cat << EOF > "$FILE_STASH_URL_DIR/delete.sql"
-- name: DeleteFileStashURLByID :exec
DELETE FROM file_stash_url WHERE id = \$1;
EOF

# RabbitMQ Config
cat << EOF > "$RABBIT_MQ_CONFIG_DIR/create.sql"
-- name: CreateRabbitMQConfig :exec
INSERT INTO rabbit_mq_config (conn_url, username, password) VALUES (\$1, \$2, \$3);
EOF

cat << EOF > "$RABBIT_MQ_CONFIG_DIR/read.sql"
-- name: GetRabbitMQConfig :one
SELECT conn_url, username, password, created_at FROM rabbit_mq_config;
EOF

cat << EOF > "$RABBIT_MQ_CONFIG_DIR/update.sql"
-- name: UpdateRabbitMQConfig :exec
UPDATE rabbit_mq_config SET conn_url = \$1, username = \$2, password = \$3;
EOF

cat << EOF > "$RABBIT_MQ_CONFIG_DIR/delete.sql"
-- name: DeleteRabbitMQConfig :exec
DELETE FROM rabbit_mq_config;
EOF

echo "SQL files created successfully!"
