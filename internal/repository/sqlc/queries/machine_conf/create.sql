-- name: CreateMachineConf :exec
INSERT INTO machine_conf (machine_id, hostname, username, passphrase, port_number, password, host_key, folder_path) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
