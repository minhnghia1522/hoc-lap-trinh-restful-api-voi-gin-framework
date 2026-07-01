package app

import (
	v1handler "user-management-api/internal/handler/v1"
	"user-management-api/internal/repository"
	"user-management-api/internal/routes"
	v1routes "user-management-api/internal/routes/v1"
	v1service "user-management-api/internal/service/v1"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule(ctx *ModuleContext) *UserModule {
	// Initialize repository
	userRepo := repository.NewUserRepository(ctx.pool)
	// Initialize service
	userService := v1service.NewUserService(userRepo, ctx.Cache)
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
