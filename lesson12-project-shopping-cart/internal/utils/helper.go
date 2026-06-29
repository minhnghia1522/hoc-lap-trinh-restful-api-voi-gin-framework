package utils

import (
	"os"
	"user-management-api/pkg/logger"

	"github.com/rs/zerolog"
)

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func NewLoggerWithPath(path string, level string) *zerolog.Logger {
	config := logger.LoggerConfig{
		Level:       level,
		FileName:    path,
		MaxSize:     1, // megabytes
		MaxBackups:  5,
		MaxAge:      5,    //days
		Compress:    true, // disabled by default
		Environment: GetEnv("APP_ENV", "development"),
	}
	return logger.NewLogger(config)
}
