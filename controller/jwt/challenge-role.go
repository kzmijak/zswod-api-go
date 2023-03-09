package jwtController

import (
	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/models/role"
	"golang.org/x/exp/slices"
)

func (c JwtController) challengeRole(userRole role.Role, ctx *gin.Context) bool {
	tokenString := utils.ExtractToken(ctx)
	tokenRole, err := c.JwtService.ExtractTokenRole(tokenString)

	if err != nil || !slices.Contains(userRole.OrHigher(), *tokenRole) {
		return false
	}

	return true
}