package main

import (
	"path/filepath"
	"user-management-api/internal/config"
	"user-management-api/internal/utils"
	"user-management-api/pkg/logger"

	"github.com/joho/godotenv"
)

func NewWorker(cfg *config.Config) {

}

func main() {
	rootDir := utils.MustGetWorkingDir()
	logFile := filepath.Join(rootDir, "internal/logs/app.log")
	logger.InitLogger(logger.LoggerConfig{
		Level:       "info",
		FileName:    logFile,
		MaxSize:     1, // megabytes
		MaxBackups:  5,
		MaxAge:      5,    //days
		Compress:    true, // disabled by default
		Environment: utils.GetEnv("APP_ENV", "development"),
	})

	if err := godotenv.Load(filepath.Join(rootDir, ".env")); err != nil {
		logger.Log.Warn().Msg("⚠️ No .env file found")
	} else {
		logger.Log.Info().Msg("✅ Loaded .env successfully")
	}
	// Initialize application
	appConfig := config.NewConfig()
	NewWorker(appConfig)
}
