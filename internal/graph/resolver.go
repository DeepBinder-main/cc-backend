package graph

import (
	// "github.com/Deepbinder-main/cc-backend/internal/repository"
	"github.com/Deepbinder-main/cc-backend/internal/repository"
	"github.com/jmoiron/sqlx"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB   *sqlx.DB
	Repo *repository.JobRepository
}
