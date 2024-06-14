package services

import (
	"github.com/google/uuid"
	"src/internal/repositories"
	"src/internal/schemes"
)

type UserService struct {
	repo repositories.User
}

func NewUserService(repo User) User { return &UserService{repo: repo} }

func (srv *UserService) CreateUser(user schemes.UserRequest) (schemes.UserResponse, error) {
	return srv.repo.CreateUser(user)
}

func (srv *UserService) GetUserByUUID(userUUID uuid.UUID) (schemes.UserResponse, error) {
	return srv.repo.GetUserByUUID(userUUID)
}
