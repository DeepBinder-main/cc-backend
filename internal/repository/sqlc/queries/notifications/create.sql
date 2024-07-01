-- name: CreateNotification :exec
INSERT INTO notifications (message) VALUES ($1);
