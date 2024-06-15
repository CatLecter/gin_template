package services

import (
	"github.com/gin-gonic/gin"
	"src/internal/repositories"
)

type UserServiceInterface interface {
	CreateUser(ctx *gin.Context)
	GetUserByUUID(ctx *gin.Context)
}

type Service struct{ UserServiceInterface }

func NewService(repos *repositories.Repository) *Service {
	return &Service{UserServiceInterface: NewUserService(repos.UserRepositoryInterface)}
}
