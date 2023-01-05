package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/models/role"
	"github.com/kzmijak/zswod_api_go/services/user"
	"golang.org/x/exp/slices"
)

const (
	ErrPasswordsNotMatching = "ErrPasswordsNotMatching: Passwords should match"
	ErrInvalidRole = "ErrInvalidRole: User has invalid role"
	ErrForbiddenRole = "ErrForbiddenRole: Only website administrators can create users with this role"
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

	elevatedRoles := []role.Role{
		role.Admin, 
		role.Teacher,
	}

	requestRole, _ := role.FromString(requestBody.Role)
	if slices.Contains(elevatedRoles, requestRole) {
		tokenRole, err := c.ExtractTokenRole(ctx)

		if err != nil || *tokenRole != role.Admin  {
			ctx.JSON(http.StatusForbidden, ErrForbiddenRole)
			return 
		}
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
