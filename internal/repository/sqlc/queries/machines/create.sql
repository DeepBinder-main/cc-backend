-- name: CreateMachine :exec
INSERT INTO machines (machine_id, hostname, os_version, ip_address) VALUES ($1, $2, $3, $4);
