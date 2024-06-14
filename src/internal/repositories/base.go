package repositories

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"src/internal/schemes"
)

type User interface {
	CreateUser(user schemes.UserRequest) (schemes.UserResponse, error)
	GetUserByUUID(userUUID uuid.UUID) (schemes.UserResponse, error)
}

type Repository struct{ User }

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
	}
}
