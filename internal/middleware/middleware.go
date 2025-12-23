package middleware

import "github.com/gin-gonic/gin"

func Logger() gin.HandlerFunc {
	return gin.Logger()
}

func Recovery() gin.HandlerFunc {
	return gin.Recovery()
}
