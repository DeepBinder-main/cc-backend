-- name: GetLogicalVolumeByID :one
SELECT lv_id, machine_id, lv_name, vg_name, lv_attr, lv_size FROM logical_volumes WHERE lv_id = $1;

-- name: ListLogicalVolumes :many
SELECT lv_id, machine_id, lv_name, vg_name, lv_attr, lv_size FROM logical_volumes ORDER BY lv_id;
