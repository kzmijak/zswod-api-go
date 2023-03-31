package blobController

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/models/blobModel"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/samber/lo"
)

type BlobResponse struct {
	Title     string    `json:"title"`
	Alt       string    `json:"alt"`
	Id        string    `json:"id"`
	CreateTime time.Time `json:"createTime"`
}

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

	blobUrls := lo.Map(response.Blobs, func(item blobModel.BlobModel, index int) BlobResponse {
		return BlobResponse{
			Id: item.ID.String(),
			CreateTime: item.CreateTime,
		}
	})

	
	ctx.IndentedJSON(http.StatusOK, GetBlobsListResponse {
		Blobs: blobUrls,
		Eof: response.Eof,
	})
}