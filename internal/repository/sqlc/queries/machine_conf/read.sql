-- name: GetMachineConfByID :one
SELECT id, machine_id, hostname, username, passphrase, port_number, password, host_key, folder_path FROM machine_conf WHERE id = $1;

-- name: ListMachineConfs :many
SELECT id, machine_id, hostname, username, passphrase, port_number, password, host_key, folder_path FROM machine_conf ORDER BY id;
