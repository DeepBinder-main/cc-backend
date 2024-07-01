-- name: GetRabbitMQConfig :one
SELECT conn_url, username, password, created_at FROM rabbit_mq_config;
