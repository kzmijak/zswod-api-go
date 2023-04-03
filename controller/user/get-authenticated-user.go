package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/database"
)

func (c UserController) GetAuthenticatedUser(ctx *gin.Context) {
	var err error
	defer utils.HandleError(&err, ctx)
	token := utils.ExtractToken(ctx)

	id, err := c.JwtService.ExtractTokenID(token)
	if err != nil {
		return
	}

	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()

	user, err := c.UserService.GetUserById(id, tx)
	if err != nil {
		return 
	}

	ctx.JSON(http.StatusOK, user)
}