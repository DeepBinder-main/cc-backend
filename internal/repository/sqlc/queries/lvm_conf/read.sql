-- name: GetLVMConfByID :one
SELECT id,machine_id , username, minAvailableSpaceGB, maxAvailableSpaceGB FROM lvm_conf WHERE id = $1;

-- name: ListLVMConfs :many
SELECT id, machine_id, username, minAvailableSpaceGB, maxAvailableSpaceGB FROM lvm_conf ORDER BY id;
