-- name: GetInfluxDBConfigurationByID :one
SELECT * FROM influxdb_configurations WHERE id = $1;

-- name: ListInfluxDBConfigurations :many
SELECT * FROM influxdb_configurations ORDER BY id;
