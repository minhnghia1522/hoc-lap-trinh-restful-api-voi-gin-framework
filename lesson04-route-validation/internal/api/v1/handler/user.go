package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"nghiadev.con/hoc-golang/utils"
)

type UserHandler struct {
}

type GetUserByIdV1Param struct {
	ID int `uri:"id" binding:"gt=0"`
}

type GetUserByUuIdV1Param struct {
	Uuid string `uri:"uuid" binding:"uuid"`
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all users (V1)",
		"data": []string{
			"user 1",
			"user 2",
			"user 3",
		},
	})
}

func (u *UserHandler) GetUserByIdV1(ctx *gin.Context) {
	var params GetUserByIdV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID (V1)",
		"data": gin.H{
			"user_id": params.ID,
		},
	})
}

func (u *UserHandler) GetUserByUuIdV1(ctx *gin.Context) {
	var params GetUserByUuIdV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by UUID (V1)",
		"data": gin.H{
			"user_id": params.Uuid,
		},
	})
}

func (u *UserHandler) PostUsersV1(ctx *gin.Context) {
	// userId := ctx.Request.Body("")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Created user successful (V1)",
		// "data": gin.H{
		// 	"user_id": userId,
		// },
	})
}

func (u *UserHandler) PutUsersV1(ctx *gin.Context) {
	userId := ctx.Param("id")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Updated user successful (V1)",
		"data": gin.H{
			"user_id": userId,
		},
	})
}

func (u *UserHandler) DeleteUsersV1(ctx *gin.Context) {
	// userId := ctx.Param("id")
	ctx.JSON(http.StatusNoContent, nil)
}
