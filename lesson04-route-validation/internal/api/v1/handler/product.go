package v1handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"nghiadev.con/hoc-golang/utils"
)

type ProductHandler struct {
}

type GetProductBySlugV1Param struct {
	Slug string `uri:"slug" binding:"slug,min=3,max=100"`
}

type GetProductsV1Param struct {
	Search string `form:"search" binding:"required,min=1,max=300,search"`
}

type PostProductsV1Param struct {
	Name         string       `json:"name" binding:"required,min=3,max=100"`
	Price        int          `json:"price" binding:"required,min=100"`
	Display      bool         `json:"display" binding:"omitempty"`
	ProductImage ProductImage `json:"product_image" binding:"required"`
}

type ProductImage struct {
	ImageName string `json:"image_name" binding:"required"`
	ImageLink string `json:"image_link" binding:"required"`
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (*ProductHandler) GetProductsV1(ctx *gin.Context) {
	var params GetProductsV1Param
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	limitStr := ctx.DefaultQuery("limit", "10")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Limit must be a positive number",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all products (V1)",
		"meta": gin.H{
			"limit":      limit,
			"key-search": params.Search,
		},
		"data": []string{
			"product 1",
			"product 2",
			"product 3",
		},
	})
}

func (*ProductHandler) GetProductByIdV1(ctx *gin.Context) {
	productId := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get product by ID (V1)",
		"data": gin.H{
			"product_id": productId,
		},
	})
}

func (*ProductHandler) GetProductBySlugV1(ctx *gin.Context) {
	var params GetProductBySlugV1Param
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get product by Slug (V1)",
		"data": gin.H{
			"slug": params.Slug,
		},
	})
}

func (*ProductHandler) PostProductsV1(ctx *gin.Context) {
	var params PostProductsV1Param
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	// productId := ctx.Request.Body("")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Created product successful (V1)",
		"data": gin.H{
			"name":          params.Name,
			"price":         params.Price,
			"display":       params.Display,
			"product_image": params.ProductImage,
		},
	})
}

func (*ProductHandler) PutProductsV1(ctx *gin.Context) {
	productId := ctx.Param("id")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Updated product successful (V1)",
		"data": gin.H{
			"product_id": productId,
		},
	})
}

func (*ProductHandler) DeleteProductsV1(ctx *gin.Context) {
	// productId := ctx.Param("id")
	ctx.JSON(http.StatusNoContent, nil)
}
