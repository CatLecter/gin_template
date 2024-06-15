package repositories

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"src/internal/schemes"
)

type UserRepositoryInterface interface {
	CreateUser(user *schemes.UserRequest) (*schemes.UserResponse, error)
	GetUserByUUID(userUUID *uuid.UUID) (*schemes.UserResponse, error)
	CheckUserByPhone(phone *string) (*bool, error)
}

type Repository struct{ UserRepositoryInterface }

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		UserRepositoryInterface: NewUserRepository(db),
	}
}
