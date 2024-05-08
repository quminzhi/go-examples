package middlewares

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"matt": "123456",
	})
}
