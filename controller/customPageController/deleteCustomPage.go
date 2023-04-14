package customPageController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/utils/parser"
)

func (c CustomPageController) DeleteCustomPage(ctx *gin.Context) {
	var err error
	var status = http.StatusBadRequest
	defer utils.HandleError(&err, ctx, &status)
	
	customPageIdString := ctx.Param("id")
	customPageId, err := parser.ParseUuid(customPageIdString)
	if err != nil {
		return
	}
	
	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()

	err = c.CustomPageService.DeleteCustomPageById(customPageId, tx)
	if err != nil {
		return
	}

	tx.Commit()
	ctx.JSON(http.StatusAccepted, nil)
}