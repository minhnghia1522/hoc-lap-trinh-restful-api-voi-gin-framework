package routes

import (
	"user-management-api/internal/middleware"
	"user-management-api/internal/utils"
	"user-management-api/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Route interface {
	Register(r *gin.RouterGroup)
}

func RegisterRoutes(r *gin.Engine, routes ...Route) {
	httpLogger := newLoggerWithPath("../../internal/logs/recovery.log", "info")
	recoveryLogger := newLoggerWithPath("../../internal/logs/recovery.log", "warning")

	r.Use(
		middleware.RateLimiterMiddleware(),
		middleware.LoggerMiddleware(httpLogger),
		middleware.RecoveryMiddleware(recoveryLogger),
		middleware.APIKeyMiddleware(),
		middleware.AuthMiddleware(),
	)

	api := r.Group("/api/v1")

	for _, route := range routes {
		route.Register(api)
	}
}

func newLoggerWithPath(path string, level string) *zerolog.Logger {
	config := logger.LoggerConfig{
		Level:       level,
		FileName:    path,
		MaxSize:     1, // megabytes
		MaxBackups:  5,
		MaxAge:      5,    //days
		Compress:    true, // disabled by default
		Environment: utils.GetEnv("APP_ENV", "development"),
	}
	return logger.NewLogger(config)
}
