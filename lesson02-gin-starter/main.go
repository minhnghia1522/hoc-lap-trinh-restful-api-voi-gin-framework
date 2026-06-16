package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": []string{"Alice", "Bob", "Charlie"},
		})
	})

	r.GET("/user/:user_id", func(ctx *gin.Context) {
		userId := ctx.Param("user_id")
		ctx.JSON(200, gin.H{
			"data": "User ID: " + userId,
		})
	})

	r.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": []string{"Product 1", "Product 2", "Product 3"},
		})
	})

	r.GET("/product/:product_name", func(ctx *gin.Context) {
		productName := ctx.Param("product_name")

		price := ctx.Query("price")
		color := ctx.Query("color")
		ctx.JSON(200, gin.H{
			"data": gin.H{
				"product_name": productName,
				"price":        price,
				"color":        color,
			},
		})
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
