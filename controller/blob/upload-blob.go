package blobController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrInvalidFile = "ErrInvalidFile: Uploaded file is invalid"
)

func (c *BlobController) UploadBlob(ctx *gin.Context) {
	c.Log.Trace("Uploading blob")
	var err error
	defer utils.HandleError(&err, ctx)
	
	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()
	
	form, err := ctx.MultipartForm()

	files := form.File["files[]"]
	if err != nil {
		err = errors.Error(ErrInvalidFile)
		return
	}

	var blobs []BlobResponse

	for _, file := range files {
		response, err := c.BlobService.StoreBlob(file, tx)
		if err != nil {
			return
		}

		blobs = append(blobs, BlobResponse{
			Id: response.ID.String(),
			CreateTime: response.CreateTime,
		})
	}

	tx.Commit()
	ctx.IndentedJSON(http.StatusOK, blobs)
}