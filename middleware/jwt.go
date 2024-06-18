package middleware

import (
	"capstone-project/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		authHeader, err := ctx.Cookie("session_token")
		if err != nil || authHeader == "" {
			if ctx.Request.Header.Get("Content-Type") == "application/json" {
				ctx.JSON(http.StatusUnauthorized, model.NewErrorResponse(http.StatusUnauthorized, "Unauthorized"))
				ctx.Abort()
				return
			}
			ctx.Redirect(http.StatusSeeOther, "/login")
			return
		}

		claims := &model.Claims{}

		token, err := jwt.ParseWithClaims(authHeader, claims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, model.NewErrorResponse(http.StatusUnauthorized, "Unauthorized"))
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Bad Request"))
			ctx.Abort()
			return
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, model.NewErrorResponse(http.StatusUnauthorized, "Unauthorized"))
			ctx.Abort()
			return
		}

		ctx.Set("username", claims.Username)
		ctx.Next()
	})
}
