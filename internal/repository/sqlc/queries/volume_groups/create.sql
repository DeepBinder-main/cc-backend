-- name: CreateVolumeGroup :exec
INSERT INTO volume_groups (machine_id, vg_name, pv_count, lv_count, snap_count, vg_attr, vg_size, vg_free) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
