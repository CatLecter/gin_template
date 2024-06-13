package handlers

import (
	"github.com/gin-gonic/gin"
	"src/internal/services"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	v1 := router.Group("/api/v1")
	user := v1.Group("/user")
	{
		user.POST("/", h.CreateUser)
	}
	return router
}
