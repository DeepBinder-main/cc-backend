-- name: DeleteVolumeGroupByID :exec
DELETE FROM volume_groups WHERE vg_id = $1;
