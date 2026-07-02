package db

import (
	"context"
	"fmt"
	"time"
	"user-management-api/internal/config"
	"user-management-api/internal/db/sqlc"
	"user-management-api/internal/utils"
	"user-management-api/pkg/logger"
	"user-management-api/pkg/pgx"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
)

var (
	Pool *pgxpool.Pool
	DB   *sqlc.Queries
)

func InitDB(appConfig *config.Config) error {
	connStr := appConfig.DNS()
	var err error
	conf, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return fmt.Errorf("error parsing config %v", err)
	}

	sqlLogger := utils.NewLoggerWithPath("sql.log", "info")

	conf.MaxConns = 50
	conf.MinConns = 5
	conf.MaxConnLifetime = 30 * time.Minute
	conf.MaxConnIdleTime = 5 * time.Minute
	conf.HealthCheckPeriod = 1 * time.Minute

	conf.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger: &pgx.PgxZerologTracer{
			Logger:         *sqlLogger,
			SlowQueryLimit: 500 * time.Millisecond,
		},
		LogLevel: tracelog.LogLevelDebug,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	DBPool, err := pgxpool.NewWithConfig(ctx, conf)

	if err != nil {
		return fmt.Errorf("failed to connect database %v", err)

	}

	if err := DBPool.Ping(ctx); err != nil {
		DBPool.Close()
		return fmt.Errorf("DB ping error %v", err)
	}

	Pool = DBPool
	DB = sqlc.New(Pool)

	logger.Log.Info().Msg("Database connected")
	return nil
}
