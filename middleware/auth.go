package middleware

import (
	"net/http"

	"github.com/Milkado/api-ars-arcanum/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "No token provided"})
			return
		}

		err := helpers.ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		ctx.Next()
	}
}
