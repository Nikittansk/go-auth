package middlewares

import (
	"net/http"

	"github.com/Nikittansk/go-auth/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "request does not contain an access token",
			})
			ctx.Abort()
			return
		}

		if err := auth.ValidateToken(tokenString); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
