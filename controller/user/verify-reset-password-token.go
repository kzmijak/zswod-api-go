package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/utils/parser"
)

func (c UserController) VerifyResetPasswordToken(ctx *gin.Context) {
	tokenString := ctx.Query("token")
	var err error
	defer utils.HandleError(&err, ctx)

	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()

	tokenId, err := parser.ParseUuid(tokenString)
	if err != nil {
		return
	}
	_, err = c.ResetPasswordTokenService.GetResetPasswordTokenById(tokenId, tx)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, true)
}
