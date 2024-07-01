-- name: DeleteRealtimeLogByID :exec
DELETE FROM realtime_logs WHERE id = $1;
