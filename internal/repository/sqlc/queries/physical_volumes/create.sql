-- name: CreatePhysicalVolume :exec
INSERT INTO physical_volumes (machine_id, pv_name, vg_name, pv_fmt, pv_attr, pv_size, pv_free) VALUES ($1, $2, $3, $4, $5, $6, $7);
