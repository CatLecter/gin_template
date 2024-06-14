package services

import (
	"github.com/google/uuid"
	"src/internal/repositories"
	"src/internal/schemes"
)

type User interface {
	CreateUser(user schemes.UserRequest) (schemes.UserResponse, error)
	GetUserByUUID(userUUID uuid.UUID) (schemes.UserResponse, error)
}

type Service struct{ User }

func NewService(repos *repositories.Repository) *Service {
	return &Service{User: NewUserService(repos.User)}
}
