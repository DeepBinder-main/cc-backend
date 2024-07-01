package repository 

import (
	"github.com/Deepbinder-main/cc-backend/internal/repository/sqlc/db/db.go"
)

// Service define a service
type Service struct {
	r *db.Queries
}