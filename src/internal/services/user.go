package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"src/internal/repositories"
	"src/internal/schemes"
	"src/internal/utils"
)

type UserService struct {
	repo repositories.UserRepositoryInterface
}

func NewUserService(repo repositories.UserRepositoryInterface) UserServiceInterface {
	return &UserService{repo: repo}
}

func (srv *UserService) CreateUser(ctx *gin.Context) {
	user := schemes.UserRequest{}
	if err := ctx.BindJSON(&user); err != nil {
		log.Errorf("Error parsing body: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.NewHTTPError(fmt.Sprintf("cannot parse JSON")))
		return
	}
	isExists, err := srv.repo.CheckUserByPhone(&user.Phone)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.NewHTTPError(fmt.Sprintf("cannot check user")))
		return
	}
	if *isExists {
		ctx.AbortWithStatusJSON(
			http.StatusUnprocessableEntity,
			utils.NewHTTPError(fmt.Sprintf("user with phone number %s already exists", user.Phone)),
		)
		return
	}
	userResp, err := srv.repo.CreateUser(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.NewHTTPError("user creation error"))
		return
	}
	ctx.JSON(http.StatusOK, userResp)
	return
}

func (srv *UserService) GetUserByUUID(ctx *gin.Context) {
	userUUID, err := uuid.Parse(ctx.Param("uuid"))
	if err != nil {
		log.Errorf("Error parsing user UUID: %v", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.NewHTTPError("cannot parse user UUID"))
		return
	}
	userResp, err := srv.repo.GetUserByUUID(&userUUID)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusNotFound, utils.NewHTTPError(fmt.Sprintf("the user with UUID=%v was not found", userUUID)),
		)
		return
	}
	ctx.JSON(http.StatusOK, userResp)
	return
}
