package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/services/user"
)

func (c *Controller) GetAllUsers(ctx *gin.Context) {
	c.log.Trace("Getting users")
	users, err := c.userService.GetAllUsers()

	if err != nil {
		c.log.Error(err)
	}

	ctx.IndentedJSON(http.StatusOK, users)
}

func (c Controller) CreateUser(ctx *gin.Context) {
	var requestBody user.CreateUserRequest 

	c.log.Trace("Creating user")
	
	if err := ctx.BindJSON(&requestBody); err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, err)
	}

	response, err := c.userService.CreateUser(requestBody)

	c.log.Trace("Response Email: " + response.Username)

	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, err)
	}

	c.log.Trace("Creating user success")

	ctx.IndentedJSON(http.StatusOK, response)
}

func (c *Controller) SignIn(ctx *gin.Context) {
	var requestBody user.SignInRequest

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return;
	}

	response, err := c.userService.SignIn(requestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return;
	}

	ctx.IndentedJSON(http.StatusOK, response)
}