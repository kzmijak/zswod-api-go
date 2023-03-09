package jwtController

import "github.com/gin-gonic/gin"

func (c JwtController) throwOnErr(status *int, ctx *gin.Context) {
	if *status != 0 {
		ctx.JSON(*status, *status)
		ctx.Abort()
	}
}