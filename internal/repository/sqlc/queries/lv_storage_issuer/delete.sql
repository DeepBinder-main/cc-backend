-- name: DeleteLVStorageIssuerByID :exec
DELETE FROM lv_storage_issuer WHERE id = $1;
