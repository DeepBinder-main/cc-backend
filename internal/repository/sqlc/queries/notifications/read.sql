-- name: GetNotificationByID :one
SELECT id, message, created_at FROM notifications WHERE id = $1;

-- name: ListNotifications :many
SELECT id, message, created_at FROM notifications ORDER BY created_at DESC;
