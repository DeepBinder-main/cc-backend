-- name: DeleteMachineByID :exec
DELETE FROM machines WHERE machine_id = $1;
