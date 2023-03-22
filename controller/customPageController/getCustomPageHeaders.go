package customPageController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/database"
)

func (c CustomPageController) GetCustomPageHeaders(ctx *gin.Context) {
	var err error
	defer utils.HandleError(&err, ctx)
	
	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()

	headers, err := c.CustomPageService.GetCustomPageHeaders(tx)
	if err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusOK, headers)
}