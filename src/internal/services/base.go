package services

import (
	"src/internal/repositories"
	"src/internal/schemes"
)

type User interface {
	CreateUser(user schemes.UserRequest) (schemes.UserResponse, error)
}

type Service struct{ User }

func NewService(repos *repositories.Repository) *Service {
	return &Service{User: NewUserService(repos.User)}
}
