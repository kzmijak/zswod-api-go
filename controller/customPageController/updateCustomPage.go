package customPageController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/models/customPageModel"
	"github.com/kzmijak/zswod_api_go/modules/database"
)

type UpdateCustomPageRequest struct {
	Id uuid.UUID `json:"id"`
	CustomPage customPageModel.UpdateCustomPagePayload `json:"customPage"`
}

func (c CustomPageController) UpdateCustomPage(ctx *gin.Context) {
	var requestBody UpdateCustomPageRequest
	var err error
	var status = http.StatusBadRequest
	defer utils.HandleError(&err, ctx, &status)
	
	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()
	
	if err = ctx.BindJSON(&requestBody); err != nil {
		return
	}

	customPageUpdated, err := c.CustomPageService.UpdateCustomPage(requestBody.Id, requestBody.CustomPage, tx)
	if err != nil {
		return
	}

	tx.Commit()
	ctx.IndentedJSON(http.StatusOK, customPageUpdated)
}