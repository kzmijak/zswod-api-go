package jwtController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/models/role"
)

func (c JwtController) RequireTeacher(ctx *gin.Context) {
	var status int
	defer c.throwOnErr(&status, ctx)

	token := utils.ExtractToken(ctx)
	if !c.JwtService.CheckValid(token) {
		status = http.StatusUnauthorized
		return
	}

	if satisfied := c.challengeRole(role.Teacher, ctx); !satisfied {
		status = http.StatusForbidden
		return
	}
	ctx.Next()
}