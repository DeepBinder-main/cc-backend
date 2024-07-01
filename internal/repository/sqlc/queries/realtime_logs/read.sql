-- name: GetRealtimeLogByID :one
SELECT id,machine_id ,log_message, created_at FROM realtime_logs WHERE id = $1;

-- name: ListRealtimeLogs :many
SELECT id,machine_id ,log_message, created_at FROM realtime_logs ORDER BY created_at DESC;
