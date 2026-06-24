package handler

import (
	"net/http"
	"user-management-api/internal/dto"
	"user-management-api/internal/service"
	"user-management-api/internal/utils"
	"user-management-api/internal/validation"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

type GetUserByUUIDParam struct {
	Uuid string `uri:"uuid" binding:"uuid"`
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
	var params GetUserByUUIDParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseBadRequest(ctx, validation.HandleValidationErrors(err))
		return
	}

	model, err := uh.service.FindUserByUUID(params.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, model)
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var createUserRequest dto.CreateUserInput
	if err := ctx.ShouldBindBodyWithJSON(&createUserRequest); err != nil {
		utils.ResponseBadRequest(ctx, validation.HandleValidationErrors(err))
		return
	}
	userCreated, err := uh.service.CreateUser(createUserRequest)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusCreated, dto.MapUserToDTO(userCreated))
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var updateUserRequest dto.UpdateUserInput
	if err := ctx.ShouldBindBodyWithJSON(&updateUserRequest); err != nil {
		utils.ResponseBadRequest(ctx, validation.HandleValidationErrors(err))
		return
	}
	uh.service.UpdateUser(updateUserRequest)
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("uuid")
	uh.service.DeleteUser(userId)
}
