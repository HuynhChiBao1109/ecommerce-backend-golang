package middlewares

import (
	"ecommerce-backend-golang/pkg/response"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.ErrorReponse(c, response.ErrUnauthorized, "")
		}
		c.Next()
	}
}
