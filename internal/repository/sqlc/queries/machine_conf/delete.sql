-- name: DeleteMachineConfByID :exec
DELETE FROM machine_conf WHERE id = $1;
