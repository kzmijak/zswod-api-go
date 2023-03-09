package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

func (c UserController) VerifyResetPasswordToken(ctx *gin.Context) {
	tokenString := ctx.Query("token")
	var err error
	defer utils.HandleError(&err, ctx)

	tokenUuid, err := uuid.Parse(tokenString)
	if err != nil {
		err = errors.Error(ErrCouldNotParseUuid)
		return
	}

	_, err = c.UserService.GetResetPasswordTokenOwner(tokenUuid)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, true)
}
