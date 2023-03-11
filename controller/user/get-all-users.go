package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/database"
)

func (c UserController) GetAllUsers(ctx *gin.Context) {
	var err error
	var status = http.StatusBadRequest
	defer utils.HandleError(&err, ctx, &status)

	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()


	c.Log.Trace("Getting users")
	users, err := c.UserService.GetAllUsers(tx)

	if err != nil {
		c.Log.Error(err)
	}

	ctx.IndentedJSON(http.StatusOK, users)
}