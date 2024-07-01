-- name: GetVolumeGroupByID :one
SELECT vg_id, machine_id, vg_name, pv_count, lv_count, snap_count, vg_attr, vg_size, vg_free FROM volume_groups WHERE vg_id = $1;

-- name: ListVolumeGroups :many
SELECT vg_id, machine_id, vg_name, pv_count, lv_count, snap_count, vg_attr, vg_size, vg_free FROM volume_groups ORDER BY vg_id;
