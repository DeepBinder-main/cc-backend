-- name: GetLVStorageIssuerByID :one
SELECT id, machine_id, hostname, username, minAvailableSpaceGB, maxAvailableSpaceGB FROM lv_storage_issuer WHERE id = $1;

-- name: ListLVStorageIssuers :many
SELECT id, machine_id, hostname, username, minAvailableSpaceGB, maxAvailableSpaceGB FROM lv_storage_issuer ORDER BY id;
