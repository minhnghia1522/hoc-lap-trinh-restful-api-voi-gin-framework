package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	server := &http.Server{
		Addr:    a.config.ServerAddress,
		Handler: a.router,
	}

	quite := make(chan os.Signal, 1)
	// syscall.SIGINT -> Ctrl C
	// syscall.SIGTERM -> Kill service
	// syscall.SIGHUP -> Reload service
	signal.Notify(quite, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		log.Printf("🍺 Server is running at %s \n", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("❌ Failed to start server. %v", err)
		}
	}()

	<-quite
	log.Println("❗Shutdown signal received ...")

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(context); err != nil {
		log.Fatalf("⛔ Server forced to shutdown ⛔. %v", err)
	}
	log.Println("🍺 Server exited gracefully.🍺")
	return nil
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
