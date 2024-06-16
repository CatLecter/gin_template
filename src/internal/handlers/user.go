package handlers

import "github.com/gin-gonic/gin"

// CreateUser   godoc
// @Summary     Creating a new user
// @Tags        User
// @Description User endpoints
// @Produce     json
// @Param       input body schemes.UserRequest true "User"
// @Success     200         {object} schemes.UserResponse
// @Failure     400,404,422 {object} utils.HTTPError
// @Failure     500         {object} utils.HTTPError
// @Failure     default     {object} utils.HTTPError
// @Router      /api/v1/user [post]
func (h *Handler) CreateUser(ctx *gin.Context) { h.service.CreateUser(ctx) }

// GetUserByUUID godoc
// @Summary      Get user by ID
// @Tags         User
// @Description  Get user by ID
// @Produce      json
// @Param        uuid path string true "User ID"
// @Success      200         {object} schemes.UserResponse
// @Failure      400,404,422 {object} utils.HTTPError
// @Failure      500         {object} utils.HTTPError
// @Failure      default     {object} utils.HTTPError
// @Router       /api/v1/user/{uuid} [get]
func (h *Handler) GetUserByUUID(ctx *gin.Context) { h.service.GetUserByUUID(ctx) }
