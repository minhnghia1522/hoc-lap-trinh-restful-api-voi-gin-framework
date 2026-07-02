package main

import (
	"log"
	"os"
	"path/filepath"
	"user-management-api/internal/app"
	"user-management-api/internal/utils"
	"user-management-api/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {
	rootDir := mustGetWorkingDir()
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

	loadEnv(filepath.Join(rootDir, ".env"))
	// Initialize application
	application := app.NewApplication()

	// Start server
	if err := application.Run(); err != nil {
		logger.Log.Error().Err(err).Msg("Application stopped with error")
		panic(err)
	}
}

func mustGetWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("❌ Unable to get working dir:", err)
	}
	return dir
}

func loadEnv(envPath string) {
	if err := godotenv.Load(envPath); err != nil {
		log.Println("⚠️ No .env file found")
		logger.Log.Warn().Err(err).Str("env_path", envPath).Msg("⚠️ No .env file found")
	} else {
		log.Println("✅ Loaded .env successfully")
		logger.Log.Info().Str("env_path", envPath).Msg("✅ Loaded .env successfully")
	}

}
