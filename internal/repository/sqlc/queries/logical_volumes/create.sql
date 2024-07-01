-- name: CreateLogicalVolume :exec
INSERT INTO logical_volumes (machine_id, lv_name, vg_name, lv_attr, lv_size) VALUES ($1, $2, $3, $4, $5);
