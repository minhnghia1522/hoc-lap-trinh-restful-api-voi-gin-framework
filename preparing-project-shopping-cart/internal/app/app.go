package app

import (
	"log"
	"user-management-api/internal/config"
	"user-management-api/internal/db"
	"user-management-api/internal/routes"
	"user-management-api/internal/validation"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Module interface {
	Routes() routes.Route
}

type Application struct {
	config  *config.Config
	router  *gin.Engine
	modules []Module
}

func NewApplication(cfg *config.Config) *Application {

	loadEnv()
	appConfig := config.NewConfig()
	db.InitDB(appConfig)

	validation.InitValidator()

	r := gin.Default()

	modules := []Module{
		NewUserModule(),
	}

	routes.RegisterRoutes(r, getModuleRoutes(modules)...)

	return &Application{
		config:  cfg,
		router:  r,
		modules: modules,
	}
}

func (a *Application) Run() error {
	return a.router.Run(a.config.ServerAddress)
}

func getModuleRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))
	for i, module := range modules {
		routeList[i] = module.Routes()
	}
	return routeList
}

func loadEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Printf("No .env file Found")
	}
}
