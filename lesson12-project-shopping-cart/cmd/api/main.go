package main

import (
	"user-management-api/internal/app"
)

func main() {
	// Initialize application
	application := app.NewApplication()

	// Start server
	if err := application.Run(); err != nil {
		panic(err)
	}
}
