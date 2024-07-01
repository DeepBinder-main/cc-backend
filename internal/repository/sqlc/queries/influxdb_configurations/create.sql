-- name: CreateInfluxDBConfiguration :exec
INSERT INTO influxdb_configurations (type, database_name, host, port, user, password, organization, ssl_enabled, batch_size, retry_interval, retry_exponential_base, max_retries, max_retry_time, meta_as_tags)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);
