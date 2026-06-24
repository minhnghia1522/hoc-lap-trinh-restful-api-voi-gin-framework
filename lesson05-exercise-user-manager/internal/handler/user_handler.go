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

type GetUsersParams struct {
	Search string `form:"search" binding:"omitempty,min=3,max=50,search"`
	Page   *int   `form:"page" binding:"omitempty,gte=1,lte=100"`
	Limit  *int   `form:"limit" binding:"omitempty,gte=1,lte=100"`
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (uh *UserHandler) Search(ctx *gin.Context) {
	var params GetUsersParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		utils.ResponseBadRequest(ctx, validation.HandleValidationErrors(err))
	}
	page := 1
	if params.Page != nil {
		page = *params.Page
	}

	limit := 10
	if params.Limit != nil {
		limit = *params.Limit
	}
	modelResultList := uh.service.Search(params.Search, page, limit)

	utils.ResponseSuccess(ctx, http.StatusOK, dto.MapUsersToDTO(modelResultList))
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
	utils.ResponseSuccess(ctx, http.StatusOK, dto.MapUserToDTO(model))
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var createUserRequest dto.CreateUserInput
	if err := ctx.ShouldBindBodyWithJSON(&createUserRequest); err != nil {
		utils.ResponseBadRequest(ctx, validation.HandleValidationErrors(err))
		return
	}
	userCreated, err := uh.service.CreateUser(createUserRequest.MapCreateInputToModel())
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusCreated, dto.MapUserToDTO(userCreated))
}

func (uh *UserHandler) UpdateUser(ctx *gin.Context) {
	var params GetUserByUUIDParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseBadRequest(ctx, validation.HandleValidationErrors(err))
		return
	}

	var updateUserRequest dto.UpdateUserInput
	if err := ctx.ShouldBindBodyWithJSON(&updateUserRequest); err != nil {
		utils.ResponseBadRequest(ctx, validation.HandleValidationErrors(err))
		return
	}
	userUpdated, err := uh.service.UpdateUser(params.Uuid, updateUserRequest.MapUpdateInputToModel())
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, dto.MapUserToDTO(userUpdated))
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
	utils.ResponseSuccess(ctx, http.StatusOK, nil)
}
