package handler

import (
	"user-management-api/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) Search(ctx *gin.Context) {
	uh.service.Search()
}

func (uh *UserHandler) GetUserByUUID(ctx *gin.Context) {
	userId := ctx.Param("uuid")
	uh.service.FindUserByUUID(userId)
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	uh.service.CreateUser()
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	uh.service.UpdateUser()
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("uuid")
	uh.service.DeleteUser(userId)
}
