package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

var validCategory = map[string]bool{
	"java":   true,
	"golang": true,
	"node":   true,
}

func (c *CategoryHandler) GetCategoriesV1(ctx *gin.Context) {
	category := ctx.Param("category")

	if !validCategory[category] {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Category must be one of: java, golang, node, python",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Category found",
	})
}
