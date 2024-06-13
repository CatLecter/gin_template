package services

import (
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
