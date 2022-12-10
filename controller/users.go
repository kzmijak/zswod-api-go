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

func (c *Controller) CreateUser(ctx *gin.Context) {
	var requestBody user.CreateUserRequest 

	c.log.Trace("Creating user")
	
	if err := ctx.BindJSON(&requestBody); err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, err)
	}

	response, err := user.CreateUser(*c.ctx, requestBody)

	c.log.Trace("Response Email: " + response.Username)

	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, err)
	}

	c.log.Trace("Creating user success")

	ctx.IndentedJSON(http.StatusOK, response)
}