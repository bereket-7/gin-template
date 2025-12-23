package router

import (
	"github.com/bereket-7/gin-template/internal/handlers"
	"github.com/bereket-7/gin-template/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	// Global middleware
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	api := r.Group("/api")
	{
		api.GET("/health", handlers.HealthCheck)
	}
}
