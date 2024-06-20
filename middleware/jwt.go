package middleware

import (
	"capstone-project/helper"
	"capstone-project/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse(http.StatusUnauthorized, "Unauthorized"))
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse(http.StatusUnauthorized, "Bearer token not found"))
			return
		}

		claims, err := helper.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.NewErrorResponse(http.StatusUnauthorized, "Invalid token"))
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
