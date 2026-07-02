package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
	"user-management-api/internal/config"
	"user-management-api/internal/db"
	"user-management-api/internal/routes"
	"user-management-api/internal/validation"
	"user-management-api/pkg/auth"
	"user-management-api/pkg/cache"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Module interface {
	Routes() routes.Route
}

type ModuleContext struct {
	pool  *pgxpool.Pool
	Cache cache.RedisCacheService
}

type Application struct {
	config  *config.Config
	router  *gin.Engine
	modules []Module
}

func NewApplication() *Application {
	loadEnv()
	appConfig := config.NewConfig()
	if err := db.InitDB(appConfig); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	redisClient := config.NewRedisClient()
	redisService := cache.NewRedisCacheService(redisClient)
	ctx := &ModuleContext{
		pool:  db.Pool,
		Cache: redisService,
	}

	validation.InitValidator()
	r := gin.Default()

	tokenService := auth.NewJWTService(redisService)

	modules := []Module{
		NewUserModule(ctx),
		NewAuthModule(ctx, tokenService, redisService, nil, nil),
	}

	routes.RegisterRoutes(r, tokenService, redisService, getModuleRoutes(modules)...)

	return &Application{
		config:  appConfig,
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
	cmd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	envPath := filepath.Join(cmd, ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Printf("No .env file Found at %s ", envPath)
	} else {
		log.Printf("Load environment from %s successful", envPath)
	}

}
