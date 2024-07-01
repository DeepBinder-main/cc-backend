CREATE TABLE
    `notifications` (
        `id` INT PRIMARY KEY AUTO_INCREMENT,
        `message` TEXT NOT NULL,
        `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
    );

CREATE TABLE
    `realtime_logs` (
        `id` INT PRIMARY KEY AUTO_INCREMENT,
        `log_message` TEXT NOT NULL,
        `machine_id` VARCHAR(255) NOT NULL,
        `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
    );

CREATE TABLE
    `lvm_conf` (
        `id` INT PRIMARY KEY AUTO_INCREMENT,
        `machine_id` VARCHAR(255) NOT NULL,
        `username` VARCHAR(255) NOT NULL,
        `minAvailableSpaceGB` FLOAT NOT NULL,
        `maxAvailableSpaceGB` FLOAT NOT NULL,
        `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
    );

CREATE TABLE
    `machines` (
        `machine_id` VARCHAR(255) PRIMARY KEY,
        `hostname` VARCHAR(255) NOT NULL,
        `os_version` VARCHAR(255) NOT NULL,
        `ip_address` VARCHAR(255) NOT NULL,
        `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
    );

CREATE TABLE
    `logical_volumes` (
        `lv_id` INT PRIMARY KEY AUTO_INCREMENT,
        `machine_id` VARCHAR(255) NOT NULL,
        `lv_name` VARCHAR(255) NOT NULL,
        `vg_name` VARCHAR(255) NOT NULL,
        `lv_attr` VARCHAR(255) NOT NULL,
        `lv_size` VARCHAR(255) NOT NULL,
        `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
    );

CREATE TABLE
    `volume_groups` (
        `vg_id` INT PRIMARY KEY AUTO_INCREMENT,
        `machine_id` VARCHAR(255) NOT NULL,
        `vg_name` VARCHAR(255) NOT NULL,
        `pv_count` VARCHAR(255) NOT NULL,
        `lv_count` VARCHAR(255) NOT NULL,
        `snap_count` VARCHAR(255) NOT NULL,
        `vg_attr` VARCHAR(255) NOT NULL,
        `vg_size` VARCHAR(255) NOT NULL,
        `vg_free` VARCHAR(255) NOT NULL,
        `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
    );

CREATE TABLE
    `physical_volumes` (
        `pv_id` INT PRIMARY KEY AUTO_INCREMENT,
        `machine_id` VARCHAR(255) NOT NULL,
        `pv_name` VARCHAR(255) NOT NULL,
        `vg_name` VARCHAR(255) NOT NULL,
        `pv_fmt` VARCHAR(255) NOT NULL,
        `pv_attr` VARCHAR(255) NOT NULL,
        `pv_size` VARCHAR(255) NOT NULL,
        `pv_free` VARCHAR(255) NOT NULL,
        `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP)
    );

CREATE TABLE
    `lv_storage_issuer` (
        `id` INT PRIMARY KEY AUTO_INCREMENT,
        `machine_id` VARCHAR(255) NOT NULL,
        `inc_buffer` INT DEFAULT 0,
        `dec_buffer` INT DEFAULT 0,
        `hostname` VARCHAR(255) NOT NULL,
        `username` VARCHAR(255) NOT NULL,
        `minAvailableSpaceGB` FLOAT NOT NULL,
        `maxAvailableSpaceGB` FLOAT NOT NULL
    );

CREATE TABLE
    `machine_conf` (
        `id` INT PRIMARY KEY AUTO_INCREMENT,
        `machine_id` VARCHAR(255) NOT NULL,
        `hostname` VARCHAR(255) NOT NULL,
        `username` VARCHAR(255) NOT NULL,
        `passphrase` LONGTEXT,
        `port_number` INT NOT NULL,
        `password` VARCHAR(255),
        `host_key` VARCHAR(255),
        `folder_path` VARCHAR(255)
    );

CREATE TABLE
    `file_stash_url` (
        `id` INT PRIMARY KEY AUTO_INCREMENT,
        `url` VARCHAR(255) NOT NULL,
        `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
        `single_row_enforcer` INT NOT NULL DEFAULT 1
    );

CREATE TABLE
    `rabbit_mq_config` (
        `conn_url` VARCHAR(255) NOT NULL,
        `username` VARCHAR(255) NOT NULL,
        `password` VARCHAR(255) NOT NULL,
        `created_at` TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
        `single_row_enforcer` INT NOT NULL DEFAULT 1
    );

CREATE TABLE
    `influxdb_configurations` (
        `id` INT PRIMARY KEY AUTO_INCREMENT,
        `type` VARCHAR(255) NOT NULL,
        `database_name` VARCHAR(255) NOT NULL,
        `host` VARCHAR(255) NOT NULL,
        `port` INT NOT NULL,
        `user` VARCHAR(255) NOT NULL,
        `password` VARCHAR(255) NOT NULL,
        `organization` VARCHAR(255) NOT NULL,
        `ssl_enabled` BOOLEAN NOT NULL,
        `batch_size` INT NOT NULL,
        `retry_interval` VARCHAR(255) NOT NULL,
        `retry_exponential_base` INT NOT NULL,
        `max_retries` INT NOT NULL,
        `max_retry_time` VARCHAR(255) NOT NULL,
        `meta_as_tags` TEXT,
        `single_row_enforcer` INT NOT NULL DEFAULT 1
    );

CREATE UNIQUE INDEX `single_row_enforcer_unique` ON `file_stash_url` (`single_row_enforcer`);

CREATE UNIQUE INDEX `single_row_enforcer_unique` ON `rabbit_mq_config` (`single_row_enforcer`);

CREATE UNIQUE INDEX `single_row_enforcer_unique` ON `influxdb_configurations` (`single_row_enforcer`);

ALTER TABLE `logical_volumes` ADD FOREIGN KEY (`machine_id`) REFERENCES `machines` (`machine_id`);

ALTER TABLE `volume_groups` ADD FOREIGN KEY (`machine_id`) REFERENCES `machines` (`machine_id`);

ALTER TABLE `physical_volumes` ADD FOREIGN KEY (`machine_id`) REFERENCES `machines` (`machine_id`);

ALTER TABLE `lv_storage_issuer` ADD FOREIGN KEY (`machine_id`) REFERENCES `machines` (`machine_id`);

ALTER TABLE `machine_conf` ADD FOREIGN KEY (`machine_id`) REFERENCES `machines` (`machine_id`);

ALTER TABLE `realtime_logs` ADD FOREIGN KEY (`machine_id`) REFERENCES `machines` (`machine_id`);

ALTER TABLE `lvm_conf` ADD FOREIGN KEY (`machine_id`) REFERENCES `machines` (`machine_id`);