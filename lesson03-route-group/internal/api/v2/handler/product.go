package v2handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (*ProductHandler) GetProductsV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all products (V2)",
		"data": []string{
			"product 1",
			"product 2",
			"product 3",
		},
	})
}

func (*ProductHandler) GetProductByIdV2(ctx *gin.Context) {
	productId := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get product by ID (V2)",
		"data": gin.H{
			"product_id": productId,
		},
	})
}

func (*ProductHandler) PostProductsV2(ctx *gin.Context) {
	// productId := ctx.Request.Body("")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Created product successful (V2)",
		// "data": gin.H{
		// 	"product_id": productId,
		// },
	})
}

func (*ProductHandler) PutProductsV2(ctx *gin.Context) {
	productId := ctx.Param("id")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Updated product successful (V2)",
		"data": gin.H{
			"product_id": productId,
		},
	})
}

func (*ProductHandler) DeleteProductsV2(ctx *gin.Context) {
	// productId := ctx.Param("id")
	ctx.JSON(http.StatusNoContent, nil)
}
