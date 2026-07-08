package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-management-api/internal/config"
	"user-management-api/internal/db"
	"user-management-api/internal/routes"
	"user-management-api/internal/utils"
	"user-management-api/internal/validation"
	"user-management-api/pkg/auth"
	"user-management-api/pkg/cache"
	"user-management-api/pkg/logger"
	"user-management-api/pkg/mail"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
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

func NewApplication(appConfig *config.Config) (*Application, error) {

	if err := db.InitDB(appConfig); err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to initialize database")
		return nil, err
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
	mailLogger := utils.NewLoggerWithPath("mail.log", "info")
	factory, err := mail.NewProviderFactory(mail.ProviderMailtrap)
	if err != nil {
		mailLogger.Error().Err(err).Msg("Failed to create mail provider factory")
		return nil, err
	}

	mailService, err := mail.NewMailService(appConfig, mailLogger, factory)
	if err != nil {
		mailLogger.Error().Err(err).Msg("Failed to initialize mail service")
		return nil, err
	}

	modules := []Module{
		NewUserModule(ctx),
		NewAuthModule(ctx, tokenService, redisService, mailService, nil),
	}

	routes.RegisterRoutes(r, tokenService, redisService, getModuleRoutes(modules)...)

	return &Application{
		config:  appConfig,
		router:  r,
		modules: modules,
	}, nil
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
		logger.Log.Info().Str("addr", server.Addr).Msg("🍺 Server is running")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Log.Fatal().Err(err).Msg("❌ Failed to start server")
		}
	}()

	<-quite
	logger.Log.Info().Msg("❗Shutdown signal received")

	context, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(context); err != nil {
		logger.Log.Fatal().Err(err).Msg("⛔ Server forced to shutdown ⛔")
	}
	logger.Log.Info().Msg("🍺 Server exited gracefully.🍺")
	return nil
}

func getModuleRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))
	for i, module := range modules {
		routeList[i] = module.Routes()
	}
	return routeList
}
