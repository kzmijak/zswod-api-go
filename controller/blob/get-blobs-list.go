package blobController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/samber/lo"
)

type GetBlobsListResponse struct {
	Blobs []BlobResponse `json:"blobs"`
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

	blobUrls := lo.Map(response.Blobs, func(item *ent.Blob, index int) BlobResponse {
		return BlobResponse{
			Title: item.Title,
			Id: item.ID.String(),
			Alt: item.Alt,
			CreateTime: item.CreateTime,
		}
	})

	
	ctx.IndentedJSON(http.StatusOK, GetBlobsListResponse {
		Blobs: blobUrls,
		Eof: response.Eof,
	})
}