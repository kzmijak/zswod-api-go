package customPageController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/database"
)

func (c CustomPageController) GetCustomPageBySectionAndTitle(ctx *gin.Context) {
	sectionString := ctx.Param("section")
	titleString := ctx.Param("title")
	url := sectionString + "/" + titleString

	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()

	response, err := c.CustomPageService.GetCustomPageByTitleNormalized(url, tx)

	if err != nil {
		utils.HandleError(&err, ctx)
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
}