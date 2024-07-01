-- name: CreateLVStorageIssuer :exec
INSERT INTO lv_storage_issuer (machine_id, hostname, username, minAvailableSpaceGB, maxAvailableSpaceGB) VALUES ($1, $2, $3, $4, $5);
