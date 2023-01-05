package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/services/user"
)

const (
	ErrPasswordsNotMatching = "ErrPasswordsNotMatching: Passwords should match"
)

func (c *Controller) GetAllUsers(ctx *gin.Context) {
	c.log.Trace("Getting users")
	users, err := c.userService.GetAllUsers()

	if err != nil {
		c.log.Error(err)
	}

	ctx.IndentedJSON(http.StatusOK, users)
}

type CreateUserRequest struct {
	user.CreateUserRequest
	PasswordConfirm string `json:"passwordConfirm"`
}

func (c Controller) CreateUser(ctx *gin.Context) {
	var requestBody CreateUserRequest 

	c.log.Trace("Creating user")
	
	if err := ctx.BindJSON(&requestBody); err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if requestBody.Password != requestBody.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, ErrPasswordsNotMatching)
		return
	}

	_, err := c.userService.CreateUser(requestBody.CreateUserRequest)

	if err != nil {
		c.log.Error(err)
		ctx.JSON(http.StatusBadRequest, err)
		return 
	}

	c.log.Trace("Creating user success")

	ctx.IndentedJSON(http.StatusAccepted, nil)
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

func (c *Controller) GetCurrentUser(ctx *gin.Context) {
	token, err := c.ExtractTokenID(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, token)

}
