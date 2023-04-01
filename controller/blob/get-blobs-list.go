package blobController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/models/blobModel"
	"github.com/kzmijak/zswod_api_go/modules/database"
)

type GetBlobsListResponse struct {
	Blobs []blobModel.BlobModel `json:"blobs"`
	Eof bool `json:"eof"`
}

func (c *BlobController) GetBlobsList(ctx *gin.Context) {
	var query utils.PaginationQuery
	var err error
	defer utils.HandleError(&err, ctx)
	
	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()


	if err = ctx.ShouldBindQuery(&query); err != nil {
		return
	}
	
	response, err := c.BlobService.GetBlobs(query.Amount, query.Offset, tx)
	if err != nil {
		return
	}

	
	ctx.IndentedJSON(http.StatusOK, GetBlobsListResponse {
		Blobs: response.Blobs,
		Eof: response.Eof,
	})
}