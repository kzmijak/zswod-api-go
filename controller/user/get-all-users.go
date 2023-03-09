package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c UserController) GetAllUsers(ctx *gin.Context) {
	c.Log.Trace("Getting users")
	users, err := c.UserService.GetAllUsers()

	if err != nil {
		c.Log.Error(err)
	}

	ctx.IndentedJSON(http.StatusOK, users)
}