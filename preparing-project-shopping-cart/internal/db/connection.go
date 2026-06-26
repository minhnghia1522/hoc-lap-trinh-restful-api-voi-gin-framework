package db

import (
	"context"
	"fmt"
	"log"
	"time"
	"user-management-api/internal/config"
	"user-management-api/internal/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *sqlc.Queries

func InitDB(appConfig *config.Config) error {
	connStr := appConfig.DNS()
	var err error
	conf, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return fmt.Errorf("error parsing config %v", err)
	}

	conf.MaxConns = 50
	conf.MinConns = 5
	conf.MaxConnLifetime = 30 * time.Minute
	conf.MaxConnIdleTime = 5 * time.Minute
	conf.HealthCheckPeriod = 1 * time.Minute

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

	DB = sqlc.New(DBPool)

	log.Println("Database connected")
	return nil
}
