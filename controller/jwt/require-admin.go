package jwtController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/models/role"
)

func (c JwtController) RequireAdmin(ctx *gin.Context) {
	var status int
	defer c.throwOnErr(&status, ctx)

	tokenString := utils.ExtractToken(ctx)
	if !c.JwtService.CheckValid(tokenString) {
		status = http.StatusUnauthorized
		return
	}

	if satisfied := c.challengeRole(role.Admin, ctx); !satisfied {
		status = http.StatusForbidden
		return
	}
	ctx.Next()
}