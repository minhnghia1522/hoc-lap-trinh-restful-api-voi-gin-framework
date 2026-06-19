package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"nghiadev.con/hoc-golang/utils"
)

type CategoryHandler struct {
}

type GetCategoriesV1Param struct {
	Category string `uri:"category" binding:"oneof=java golang node"`
}

type PostCategoriesV1Param struct {
	Name   string `form:"name" binding:"required"`
	Status string `form:"status" binding:"oneof=1 2"`
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c *CategoryHandler) GetCategoriesV1(ctx *gin.Context) {
	var params GetCategoriesV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Category found",
	})
}

func (c *CategoryHandler) PostCategoriesV1(ctx *gin.Context) {
	var params PostCategoriesV1Param
	if err := ctx.ShouldBind(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Category found",
		"data": gin.H{
			"name":   params.Name,
			"status": params.Status,
		},
	})
}
