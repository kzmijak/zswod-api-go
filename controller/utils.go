package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleError(err *error, ctx *gin.Context, status ...*int) {
	if len(status) > 1 {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	if *err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
}