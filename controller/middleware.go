package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c Controller) JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := c.ExtractToken(ctx)
		_, err := c.jwtService.ParseToken(token)
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}