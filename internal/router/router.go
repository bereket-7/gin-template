package router

import (
	"github.com/bereket-7/gin-template/internal/config"
	"github.com/bereket-7/gin-template/internal/handlers"
	"github.com/bereket-7/gin-template/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, cfg *config.Config) {
	// Global middleware
	r.Use(middleware.CORS(cfg.AllowedOrigins))
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	api := r.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)
	}
}
