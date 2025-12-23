package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bereket-7/gin-template/internal/config"
	"github.com/bereket-7/gin-template/internal/logger"
	"github.com/bereket-7/gin-template/internal/redis"
	"github.com/bereket-7/gin-template/internal/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type App struct {
	engine *gin.Engine
	config *config.Config
}

func New() *App {
	cfg := config.Load()

	redisErr := redis.Connect(cfg.Redis)
	logger.Init(cfg.Env)
	logger.Log.Info("application starting", zap.String("env", cfg.Env), zap.String("port", cfg.Port))

	if redisErr != nil {
		logger.Log.Error("failed to connect to redis", zap.Error(redisErr))
	}

	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	router.Setup(engine)

	return &App{
		engine: engine,
		config: cfg,
	}
}

func (a *App) Run() error {
	address := fmt.Sprintf(":%s", a.config.Port)

	srv := &http.Server{
		Addr:    address,
		Handler: a.engine,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("listen error", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so no need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Info("shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("server forced to shutdown", zap.Error(err))
		return err
	}

	logger.Log.Info("server exiting")
	return nil
}
