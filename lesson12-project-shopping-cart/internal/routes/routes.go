package routes

import (
	"user-management-api/internal/middleware"
	v1routes "user-management-api/internal/routes/v1"
	"user-management-api/internal/utils"
	"user-management-api/pkg/auth"
	"user-management-api/pkg/cache"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type Route interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, tokenService auth.TokenService, cacheService cache.RedisCacheService, routes ...Route) {
	httpLogger := utils.NewLoggerWithPath("../../internal/logs/http.log", "info")
	recoveryLogger := utils.NewLoggerWithPath("../../internal/logs/recovery.log", "warning")
	limiterLogger := utils.NewLoggerWithPath("../../internal/logs/limiter.log", "info")
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(
		middleware.CORSMiddleware(),
		middleware.RateLimiterMiddleware(limiterLogger),
		middleware.TraceMiddleware(),
		middleware.LoggerMiddleware(httpLogger),
		middleware.RecoveryMiddleware(recoveryLogger),
		middleware.APIKeyMiddleware(),
	)

	v1api := r.Group("/api/v1")

	protected := v1api.Group("")
	middleware.InitAuthMiddleware(tokenService, cacheService)
	protected.Use(middleware.AuthMiddleware())

	for _, route := range routes {
		switch route := route.(type) {
		case *v1routes.AuthRoutes:
			route.Register(v1api)
		default:
			route.Register(protected)
		}
	}
}
