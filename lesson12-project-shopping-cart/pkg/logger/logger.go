package logger

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

type contextKey string

const TraceIdKey contextKey = "trace_id"

type LoggerConfig struct {
	Level       string
	FileName    string
	MaxSize     int
	MaxBackups  int
	MaxAge      int
	Compress    bool
	Environment string
}

func NewLogger(config LoggerConfig) *zerolog.Logger {
	zerolog.TimeFieldFormat = time.RFC3339
	lvl, err := zerolog.ParseLevel(config.Level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(lvl)

	var writer io.Writer
	if config.Environment == "development" {
		writer = os.Stdout
	} else {
		writer = &lumberjack.Logger{
			Filename:   config.FileName,
			MaxSize:    config.MaxSize, // megabytes
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge,   //days
			Compress:   config.Compress, // disabled by default
			LocalTime:  true,
		}
	}
	logger := zerolog.New(writer).With().Timestamp().Logger()

	return &logger
}

func GetTraceID(ctx context.Context) string {
	if traceID, ok := ctx.Value(TraceIdKey).(string); ok {
		return traceID
	}

	return ""
}
