package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) CreateUser(ctx *gin.Context) { h.service.CreateUser(ctx) }

func (h *Handler) GetUserByUUID(ctx *gin.Context) { h.service.GetUserByUUID(ctx) }
