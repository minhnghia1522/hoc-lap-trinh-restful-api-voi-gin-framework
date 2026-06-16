package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (*ProductHandler) GetProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all products (V1)",
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

func (*ProductHandler) PostProductsV1(ctx *gin.Context) {
	// productId := ctx.Request.Body("")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Created product successful (V1)",
		// "data": gin.H{
		// 	"product_id": productId,
		// },
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
