-- name: UpdateRabbitMQConfig :exec
UPDATE rabbit_mq_config SET conn_url = $1, username = $2, password = $3;
