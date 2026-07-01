package routes

import (
	"user-management-api/internal/middleware"
	"user-management-api/internal/utils"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type Route interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, routes ...Route) {
	httpLogger := utils.NewLoggerWithPath("../../internal/logs/http.log", "info")
	recoveryLogger := utils.NewLoggerWithPath("../../internal/logs/recovery.log", "warning")
	limiterLogger := utils.NewLoggerWithPath("../../internal/logs/limiter.log", "info")

	r.Use(
		middleware.CORSMiddleware(),
		middleware.RateLimiterMiddleware(limiterLogger),
		middleware.TraceMiddleware(),
		middleware.LoggerMiddleware(httpLogger),
		middleware.RecoveryMiddleware(recoveryLogger),
		middleware.APIKeyMiddleware(),
		middleware.AuthMiddleware(),
	)

	r.Use(gzip.Gzip(gzip.DefaultCompression))

	api := r.Group("/api/v1")

	for _, route := range routes {
		route.Register(api)
	}
}
