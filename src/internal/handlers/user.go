package handlers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"src/internal/schemes"
)

func (h *Handler) CreateUser(c *gin.Context) {
	user := schemes.UserRequest{}
	if err := c.BindJSON(&user); err != nil {
		log.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "cannot parse JSON"})
		return
	}
	userResp, err := h.service.CreateUser(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user creation error"})
		return
	}
	c.JSON(http.StatusOK, userResp)
	return
}
