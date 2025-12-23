package app

import (
	"fmt"

	"github.com/bereket-7/gin-template/internal/config"
	"github.com/bereket-7/gin-template/internal/logger"
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
	logger.Init(cfg.Env)
	logger.Log.Info("application starting", zap.String("env", cfg.Env), zap.String("port", cfg.Port))

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
	return a.engine.Run(address)
}
