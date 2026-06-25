package db

import (
	"fmt"
	"lesson08-prepare-connection/internal/config"
)

func InitDB() error {
	connStr := config.NewConfig().DNS()

	fmt.Print("Connection string: %s", connStr)

	return nil
}
