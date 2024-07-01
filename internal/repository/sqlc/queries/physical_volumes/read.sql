-- name: GetPhysicalVolumeByID :one
SELECT pv_id, machine_id, pv_name, vg_name, pv_fmt, pv_attr, pv_size, pv_free FROM physical_volumes WHERE pv_id = $1;

-- name: ListPhysicalVolumes :many
SELECT pv_id, machine_id, pv_name, vg_name, pv_fmt, pv_attr, pv_size, pv_free FROM physical_volumes ORDER BY pv_id;
