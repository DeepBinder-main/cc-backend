-- name: CreateLVMConf :exec
INSERT INTO lvm_conf (machine_id , username, minAvailableSpaceGB, maxAvailableSpaceGB) VALUES ($1, $2, $3 , $4);
