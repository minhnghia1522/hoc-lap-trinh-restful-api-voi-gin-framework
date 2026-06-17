package main

import (
	"github.com/gin-gonic/gin"
	v1handler "nghiadev.con/hoc-golang/internal/api/v1/handler"
	v2handler "nghiadev.con/hoc-golang/internal/api/v2/handler"
	"nghiadev.con/hoc-golang/utils"
)

func main() {
	r := gin.Default()

	if err := utils.RegisterValidators(); err != nil {
		panic(err)
	}

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("users")
		{
			userHandlerV1 := v1handler.NewUserHandler()
			user.GET("/", userHandlerV1.GetUsersV1)
			user.GET("/:id", userHandlerV1.GetUserByIdV1)
			user.GET("/admin/:uuid", userHandlerV1.GetUserByUuIdV1)
			user.POST("/", userHandlerV1.PostUsersV1)
			user.PUT("/:id", userHandlerV1.PutUsersV1)
			user.DELETE("/:id", userHandlerV1.DeleteUsersV1)

		}
		product := v1.Group("products")
		productHandlerV1 := v1handler.NewProductHandler()
		product.GET("/", productHandlerV1.GetProductsV1)
		product.GET("/:slug", productHandlerV1.GetProductBySlugV1)
		product.POST("/", productHandlerV1.PostProductsV1)
		product.PUT("/:id", productHandlerV1.PutProductsV1)
		product.DELETE("/:id", productHandlerV1.DeleteProductsV1)

		category := v1.Group("categories")
		categoryHandler := v1handler.NewCategoryHandler()
		category.POST("/:category", categoryHandler.GetCategoriesV1)
	}

	v2 := r.Group("/api/v2")
	{
		user := v2.Group("users")
		{
			userHandlerV2 := v2handler.NewUserHandler()
			user.GET("/", userHandlerV2.GetUsersV2)
			user.GET("/:id", userHandlerV2.GetUserByIdV2)
			user.POST("/", userHandlerV2.PostUsersV2)
			user.PUT("/:id", userHandlerV2.PutUsersV2)
			user.DELETE("/:id", userHandlerV2.DeleteUsersV2)
		}

		product := v2.Group("products")
		productHandlerV2 := v2handler.NewProductHandler()
		product.GET("/", productHandlerV2.GetProductsV2)
		product.GET("/:id", productHandlerV2.GetProductByIdV2)
		product.POST("/", productHandlerV2.PostProductsV2)
		product.PUT("/:id", productHandlerV2.PutProductsV2)
		product.DELETE("/:id", productHandlerV2.DeleteProductsV2)
	}
	r.Run(":8080")
}
