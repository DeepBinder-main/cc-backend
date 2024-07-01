-- name: DeleteInfluxDBConfigurationByID :exec
DELETE FROM influxdb_configurations WHERE id = $1;
