package repositories

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"src/internal/schemes"
)

type User interface {
	CreateUser(user *schemes.UserRequest) (*schemes.UserResponse, error)
	GetUserByUUID(userUUID *uuid.UUID) (*schemes.UserResponse, error)
}

type Repository struct{ User }

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		User: NewUserRepository(db),
	}
}
