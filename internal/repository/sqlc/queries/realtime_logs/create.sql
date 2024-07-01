-- name: CreateRealtimeLog :exec
INSERT INTO realtime_logs (machine_id , log_message) VALUES ($1, $2);
