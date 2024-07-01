-- name: CreateRabbitMQConfig :exec
INSERT INTO rabbit_mq_config (conn_url, username, password) VALUES ($1, $2, $3);
