-- name: DeleteLogicalVolumeByID :exec
DELETE FROM logical_volumes WHERE lv_id = $1;
