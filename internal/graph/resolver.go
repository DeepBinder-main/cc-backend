package graph

import (
	sqlcdb "github.com/Deepbinder-main/cc-backend/internal/repository/sqlc/db"

	"github.com/jmoiron/sqlx"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Queries *sqlcdb.Queries
	DB      *sqlx.DB
}
