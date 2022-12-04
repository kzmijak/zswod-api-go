package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/services/user"
)

func (c *Controller) GetAllUsers(ctx *gin.Context) {
	c.log.Trace("Getting users")
	users, err := user.GetAllUsers(*c.ctx)

	if err != nil {
		c.log.Error(err)
	}

	ctx.IndentedJSON(http.StatusOK, users)
}