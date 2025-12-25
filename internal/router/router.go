package router

import (
	"github.com/bereket-7/gin-template/internal/config"
	"github.com/bereket-7/gin-template/internal/handlers"
	"github.com/bereket-7/gin-template/internal/middleware"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
