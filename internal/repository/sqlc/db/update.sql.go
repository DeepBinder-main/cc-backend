// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: update.sql

package db

import (
	"context"
)

const updateRabbitMQConfig = `-- name: UpdateRabbitMQConfig :exec
UPDATE rabbit_mq_config SET conn_url = $1, username = $2, password = $3
`

func (q *Queries) UpdateRabbitMQConfig(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, updateRabbitMQConfig)
	return err
}
