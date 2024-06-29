-- name: GetFileStashURL :one
SELECT * FROM file_stash_url LIMIT 1;

-- name: InsertFileStashURL :exec
INSERT INTO file_stash_url (url, single_row_enforcer)
VALUES (?, 1)
ON DUPLICATE KEY UPDATE url = VALUES(url);

-- name: UpdateFileStashURL :exec
UPDATE file_stash_url
SET url = ?
WHERE single_row_enforcer = 1;

-- name: DeleteFileStashURL :exec
DELETE FROM file_stash_url
WHERE single_row_enforcer = 1;

-- name: GetRabbitMQConfig :one
SELECT * FROM rabbit_mq_config LIMIT 1;

-- name: InsertRabbitMQConfig :exec
INSERT INTO rabbit_mq_config (conn_url, username, password, single_row_enforcer)
VALUES (?, ?, ?, 1)
ON DUPLICATE KEY UPDATE conn_url = VALUES(conn_url), username = VALUES(username), password = VALUES(password);

-- name: UpdateRabbitMQConfig :exec
UPDATE rabbit_mq_config
SET conn_url = ?, username = ?, password = ?
WHERE single_row_enforcer = 1;

-- name: DeleteRabbitMQConfig :exec
DELETE FROM rabbit_mq_config
WHERE single_row_enforcer = 1;