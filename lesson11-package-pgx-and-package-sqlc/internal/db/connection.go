package db

import (
	"context"
	"fmt"
	"lesson08-prepare-connection/internal/config"
	"lesson08-prepare-connection/internal/db/sqlc"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *sqlc.Queries

func InitDB() error {
	connStr := config.NewConfig().DNS()
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
