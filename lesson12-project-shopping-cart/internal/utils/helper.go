package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"user-management-api/pkg/logger"

	"github.com/rs/zerolog"
)

func GetEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func GetIntEnv(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	intVal, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return intVal
}

func NewLoggerWithPath(filename string, level string) *zerolog.Logger {
	cmd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	path := filepath.Join(cmd,"internal/logs", filename)
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

func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes), nil
}
