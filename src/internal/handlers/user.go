package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"src/internal/schemes"
)

func (h *Handler) CreateUser(ctx *gin.Context) {
	user := schemes.UserRequest{}
	if err := ctx.BindJSON(&user); err != nil {
		log.Errorf("Error parsing body: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "cannot parse JSON"})
		return
	}
	userResp, err := h.service.CreateUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user creation error"})
		return
	}
	ctx.JSON(http.StatusOK, userResp)
	return
}

func (h *Handler) GetUserByUUID(ctx *gin.Context) {
	userUUID, err := uuid.Parse(ctx.Param("uuid"))
	if err != nil {
		log.Errorf("Error parsing user UUID: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "cannot parse UUID"})
		return
	}
	userResp, err := h.service.GetUserByUUID(userUUID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error receiving the user"})
		return
	}
	ctx.JSON(http.StatusOK, userResp)
	return
}
