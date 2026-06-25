package db

import (
	"context"
	"database/sql"
	"lesson08-prepare-connection/internal/config"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() *sql.DB {
	connStr := config.NewConfig().DNS()

	var err error

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}

	db.SetMaxIdleConns(3)                   // Số kết nối tối đa
	db.SetMaxOpenConns(3)                   // Số kết nối tối đa
	db.SetConnMaxLifetime(30 * time.Minute) // Đóng kết nối sau 30 phút
	db.SetConnMaxIdleTime(5 * time.Minute)  // Đóng kết nối nhàn rỗi sau 5 phút

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		db.Close()
		log.Fatal("DB ping error", err)
	}
	log.Println("Database Connected")
	return db
}
