package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SampleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Trước khi bắt đầu vào handler (Before)

		log.Println("Start function")
		ctx.Next() // Đi vào handler
		log.Println("End function")
		// Sau khi handler xử lý xong (After)
	}
}
