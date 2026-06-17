package v1handler

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"nghiadev.con/hoc-golang/utils"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)

func (*ProductHandler) GetProductsV1(ctx *gin.Context) {
	search := ctx.Query("search")

	if err := utils.ValidationRequired("Search", search); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(search) < 3 || len(search) > 50 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Search must be between 3 and 50 characters",
		})
		return
	}

	if !searchRegex.MatchString(search) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Search must contain only letters, number and spaces",
		})
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
			"key-search": search,
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
	slug := ctx.Param("slug")

	if !slugRegex.MatchString(slug) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Slug must contain only lowercase letter, numbers, hyphens and dots.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get product by Slug (V1)",
		"data": gin.H{
			"slug": slug,
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
