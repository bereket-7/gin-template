package handlers

import (
	"net/http"

	appRedis "github.com/bereket-7/gin-template/internal/redis"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	if appRedis.Client == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "redis client not initialized",
		})
		return
	}

	err := appRedis.Client.Set(appRedis.Ctx, "health", "ok", 0).Err()

	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "redis not connected",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"redis":  "connected",
	})
}
