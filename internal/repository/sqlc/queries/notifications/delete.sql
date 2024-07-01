-- name: DeleteNotificationByID :exec
DELETE FROM notifications WHERE id = $1;
