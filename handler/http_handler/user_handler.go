package http_handler

import (
	"hacktiv8-msib-final-project-2/dto"
	"hacktiv8-msib-final-project-2/entity"
	"hacktiv8-msib-final-project-2/pkg/errs"
	"hacktiv8-msib-final-project-2/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService: userService}
}

func (u *userHandler) Register(ctx *gin.Context) {
	var requestBody dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	registeredUser, err := u.userService.Register(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, registeredUser)
}

func (u *userHandler) Login(ctx *gin.Context) {
	var requestBody dto.LoginRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	token, err := u.userService.Login(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (u *userHandler) UpdateUser(ctx *gin.Context) {
	var requestBody dto.UpdateUserRequest
	userData, ok := ctx.MustGet("userData").(*entity.User)
	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	log.Println("Email:", userData.Email)

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	updatedUser, err := u.userService.UpdateUser(userData.ID, &requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}
