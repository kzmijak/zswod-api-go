package jwtController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
)

func (c JwtController) RequireAuthenticated(ctx *gin.Context) {
	var status int
	defer c.throwOnErr(&status, ctx)

	token := utils.ExtractToken(ctx)

	if !c.JwtService.CheckValid(token) {
		status = http.StatusUnauthorized
		return
	}

	ctx.Next()
}