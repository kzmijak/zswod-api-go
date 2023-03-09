package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/services/user"
)

const (
	ErrCouldNotParseUuid = "ErrCouldNotParseUuid: Could not parse string to uuid"
)

type SetNewPasswordBody struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

func (c UserController) SetNewPassword(ctx *gin.Context) {
	var requestBody SetNewPasswordBody
	var err error
	defer utils.HandleError(&err, ctx)

	if err := ctx.BindJSON(&requestBody); err != nil {
		return
	}

	tokenUuid, err := uuid.Parse(requestBody.Token)
	if err != nil {
		err = errors.Error(ErrCouldNotParseUuid)
		return
	}

	owner, err := c.UserService.GetResetPasswordTokenOwner(tokenUuid)
	if err != nil {
		return
	}

	err = c.UserService.SetNewPassword(user.SetNewPasswordParams{
		UserId:      owner.ID,
		NewPassword: requestBody.Password,
	})
	if err != nil {
		return
	}

	err = c.UserService.InvalidateResetPasswordToken(tokenUuid)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, true)
}
