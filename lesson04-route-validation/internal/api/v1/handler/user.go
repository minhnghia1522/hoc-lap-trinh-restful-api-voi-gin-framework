package v1handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
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
	userIdStr := ctx.Param("id")
	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID must be a number",
		})
		return
	}
	if userId <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID must be a positive",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID (V1)",
		"data": gin.H{
			"user_id": userId,
		},
	})
}

func (u *UserHandler) GetUserByUuIdV1(ctx *gin.Context) {
	uuidStr := ctx.Param("uuid")
	_, err := uuid.Parse(uuidStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID must be a valid UUID",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by UUID (V1)",
		"data": gin.H{
			"user_id": uuidStr,
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
