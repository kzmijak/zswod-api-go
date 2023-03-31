package authController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/models/role"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/utils/encryption"
)

const (
	ErrUserNotFound = "ErrUserNotFound: Could not match any user with provided username"
	ErrInvalidPassword = "ErrInvalidPassword: Password for the user is incorrect"
	ErrInvalidClaims = "ErrInvalidClaims: Could not generate token for this user"
	ErrNoRoles = "ErrNoRoles: User has no roles and could not be retrieved"
	ErrUnknownRole = "ErrUnknownRole: User has unknown roles and could not be retrieved"
)


type SignInRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (c AuthController) SignIn(ctx *gin.Context) {
	var requestBody SignInRequest
	var err error
	var status = http.StatusBadRequest
	defer utils.HandleError(&err, ctx, &status)

	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()

	if err = ctx.BindJSON(&requestBody); err != nil {
		return
	}

	user, err := c.UserService.GetUserByEmail(requestBody.Email, tx)
	if err != nil {
		return 
	}

	if err = encryption.CompareWithHash(requestBody.Password, user.Password); err != nil {
		return
	}

	role, err := role.FromString(user.Role.String())
	if err != nil {
		return
	}

	jwt, err := c.JwtService.GenerateToken(string(user.ID.String()), role)

	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, jwt)
}