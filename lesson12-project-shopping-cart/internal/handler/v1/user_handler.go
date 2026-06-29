package v1handler

import (
	"net/http"
	v1dto "user-management-api/internal/dto/v1"
	v1service "user-management-api/internal/service/v1"
	"user-management-api/internal/utils"
	"user-management-api/internal/validation"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service v1service.UserService
}

type GetUserByUUIDParam struct {
	Uuid string `uri:"uuid" binding:"uuid"`
}

type GetUsersParams struct {
	Search string `form:"search" binding:"omitempty,min=3,max=50,search"`
	Page   int    `form:"page" binding:"omitempty,gte=1,lte=100"`
	Limit  int    `form:"limit" binding:"omitempty,gte=1,lte=100"`
}

func NewUserHandler(service v1service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) Search(ctx *gin.Context) {
	var params GetUsersParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		utils.ResponseBadRequest(ctx, validation.HandleValidationErrors(err))
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.Limit == 0 {
		params.Limit = 10
	}
	// modelResultList := uh.service.Search(params.Search, params.Page, params.Limit)

	utils.ResponseSuccess(ctx, http.StatusOK, "")
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
	utils.ResponseSuccess(ctx, http.StatusOK, v1dto.MapUserToDTO(model))
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var createUserRequest v1dto.CreateUserInput
	if err := ctx.ShouldBindBodyWithJSON(&createUserRequest); err != nil {
		utils.ResponseBadRequest(ctx, validation.HandleValidationErrors(err))
		return
	}
	userCreated, err := uh.service.CreateUser(ctx, createUserRequest.MapCreateInputToModel())
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusCreated, v1dto.MapUserToDTO(userCreated))
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var params GetUserByUUIDParam
	var err error
	if err = ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseBadRequest(ctx, validation.HandleValidationErrors(err))
		return
	}

	var updateUserRequest v1dto.UpdateUserInput
	if err = ctx.ShouldBindBodyWithJSON(&updateUserRequest); err != nil {
		utils.ResponseBadRequest(ctx, validation.HandleValidationErrors(err))
		return
	}
	// userUpdated, err := uh.service.UpdateUser(params.Uuid, "")
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, "")
}

func (uh *UserHandler) DeleteUser(ctx *gin.Context) {
	var params GetUserByUUIDParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseBadRequest(ctx, validation.HandleValidationErrors(err))
		return
	}
	if err := uh.service.DeleteUser(params.Uuid); err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	utils.ResponseStatusCode(ctx, http.StatusNoContent)
}
