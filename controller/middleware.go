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
	var status *int
	defer c.throwOnErr(status, ctx)

	if err != nil {
		*status = http.StatusUnauthorized
		return
	}
	
	if !c.CheckValid(ctx) {
		*status = http.StatusUnauthorized
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
	var status *int
	defer c.throwOnErr(status, ctx)

	if !c.CheckValid(ctx) {
		*status = http.StatusUnauthorized
		return
	}

	if satisfied := c.challengeRole(role.Admin, ctx); !satisfied {
		*status = http.StatusForbidden
		return
	}
	ctx.Next()
}

func (c Controller) RequireTeacher(ctx *gin.Context) {
	var status *int
	defer c.throwOnErr(status, ctx)

	if !c.CheckValid(ctx) {
		*status = http.StatusUnauthorized
		return
	}

	if satisfied := c.challengeRole(role.Teacher, ctx); !satisfied {
		*status = http.StatusForbidden
		return
	}
	ctx.Next()
}

func (c Controller) throwOnErr(status *int, ctx *gin.Context) {
	if status != nil {
		ctx.JSON(*status, *status)
		ctx.Abort()
	}
}