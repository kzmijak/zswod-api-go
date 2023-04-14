package customPageController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/models/customPageModel"
	"github.com/kzmijak/zswod_api_go/modules/database"
)

func (c CustomPageController) CreateCustomPage(ctx *gin.Context) {
	var requestBody customPageModel.CreateCustomPagePayload
	var err error
	defer utils.HandleError(&err, ctx)
	
	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()

	if err = ctx.BindJSON(&requestBody); err != nil {
		return
	}

	customPageId, err := c.CustomPageService.CreateCustomPage(requestBody, tx)
	if err != nil {
		return 
	}

	customPageUrl, err := c.CustomPageService.GetCustomPageUrlById(customPageId, tx)

	tx.Commit()
	ctx.IndentedJSON(http.StatusOK, customPageUrl)
}