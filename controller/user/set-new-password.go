package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/services/userService"
	"github.com/kzmijak/zswod_api_go/utils/parser"
)

type SetNewPasswordRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

func (c UserController) SetNewPassword(ctx *gin.Context) {
	var requestBody SetNewPasswordRequest
	var err error
	var status = http.StatusBadRequest
	defer utils.HandleError(&err, ctx, &status)

	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()

	if err = ctx.BindJSON(&requestBody); err != nil {
		return
	}

	tokenId, err := parser.ParseUuid(requestBody.Token)
	if err != nil {
		return
	}

	owner, err := c.ResetPasswordTokenService.GetResetPasswordTokenById(tokenId, tx)
	if err != nil {
		return
	}

	_, err = c.UserService.UpdateUserPassword(owner.ID.String(), userService.UpdateUserPasswordPayload{
		NewPassword: requestBody.Password,
	}, tx)
	if err != nil {
		return
	}

	err = c.ResetPasswordTokenService.InvalidateResetPasswordToken(tokenId, tx)
	if err != nil {
		return
	}
		
	tx.Commit()
	ctx.JSON(http.StatusOK, true)
}
