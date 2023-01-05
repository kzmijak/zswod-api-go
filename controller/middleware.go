package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/models/role"
	"golang.org/x/exp/slices"
)

func (c Controller) RequireAuthenticated(ctx *gin.Context)  {
	token := c.ExtractToken(ctx)
	_, err := c.jwtService.ParseToken(token)
	if err != nil {
		ctx.String(http.StatusUnauthorized, "Unauthorized")
		ctx.Abort()
		return
	}
	ctx.Next()
}

func (c Controller) challengeRole(userRole role.Role, ctx *gin.Context) bool {
	tokenRole, err := c.ExtractTokenRole(ctx)

	if err != nil || !slices.Contains(userRole.OrHigher(), *tokenRole) {
		return false
	}
	
	return true
}

func (c Controller) RequireAdmin(ctx *gin.Context) {
	if satisfied := c.challengeRole(role.Admin, ctx); !satisfied {
		ctx.JSON(http.StatusForbidden, "Forbidden")
		ctx.Abort()
		return
	}
	ctx.Next()
}

func (c Controller) RequireTeacher(ctx *gin.Context) {
	if satisfied := c.challengeRole(role.Teacher, ctx); !satisfied {
		ctx.JSON(http.StatusForbidden, "Forbidden")
		ctx.Abort()
		return
	}
	ctx.Next()
}