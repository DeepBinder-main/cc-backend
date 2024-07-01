-- name: DeleteLVMConfByID :exec
DELETE FROM lvm_conf WHERE id = $1;
