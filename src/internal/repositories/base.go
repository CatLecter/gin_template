package repositories

import (
	"github.com/jmoiron/sqlx"
	"src/internal/schemes"
)

type User interface {
	CreateUser(user schemes.UserRequest) (schemes.UserResponse, error)
}

type Repository struct{ User }

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
	}
}
