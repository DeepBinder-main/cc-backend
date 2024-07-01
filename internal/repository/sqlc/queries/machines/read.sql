-- name: GetMachineByID :one
SELECT machine_id, hostname, os_version, ip_address FROM machines WHERE machine_id = $1;

-- name: ListMachines :many
SELECT machine_id, hostname, os_version, ip_address FROM machines ORDER BY machine_id;
