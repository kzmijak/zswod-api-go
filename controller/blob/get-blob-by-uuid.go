package blobController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/utils/parser"
)

func (c *BlobController) GetBlobByUuid(ctx *gin.Context) {
	var err error
	defer utils.HandleError(&err, ctx)

	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()

	uuidString := ctx.Param("uuid")
	uuidParsed, err := parser.ParseUuid(uuidString)
	if err != nil {
		return
	}

	blob, err := c.BlobService.GetBlobById(uuidParsed, tx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.Data(http.StatusOK, blob.ContentType, blob.Blob)
}