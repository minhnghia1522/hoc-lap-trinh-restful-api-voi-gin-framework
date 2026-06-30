package app

import (
	v1handler "user-management-api/internal/handler/v1"
	"user-management-api/internal/repository"
	"user-management-api/internal/routes"
	v1routes "user-management-api/internal/routes/v1"
	v1service "user-management-api/internal/service/v1"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule(pool *pgxpool.Pool) *UserModule {
	// Initialize repository
	userRepo := repository.NewUserRepository(pool)
	// Initialize service
	userService := v1service.NewUserService(userRepo)
	// Initialize handler
	userHandler := v1handler.NewUserHandler(userService)
	// Initialize routes
	userRoutes := v1routes.NewUserRoutes(userHandler)

	return &UserModule{
		routes: userRoutes,
	}
}

func (m *UserModule) Routes() routes.Route {
	return m.routes
}
