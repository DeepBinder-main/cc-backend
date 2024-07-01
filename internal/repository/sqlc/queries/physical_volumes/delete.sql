-- name: DeletePhysicalVolumeByID :exec
DELETE FROM physical_volumes WHERE pv_id = $1;
