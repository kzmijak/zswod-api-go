package blobController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ErrInvalidFile = "ErrInvalidFile: Uploaded file is invalid"
	ErrInvalidUuidFormat = "ErrInvalidUuidFormat: Could not parse UUID. Invalid format."
)

func (c *BlobController) UploadBlob(ctx *gin.Context) {
	c.Log.Trace("Uploading blob")

	form, err := ctx.MultipartForm()

	files := form.File["files[]"]
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrInvalidFile)
		return
	}

	var blobs []BlobResponse

	for _, file := range files {
		response, err := c.BlobService.StoreBlob(file)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		blobs = append(blobs, BlobResponse{
			Title: response.Title,
			Id: response.ID.String(),
			Alt: response.Alt,
			CreatedAt: response.CreatedAt,
		})
	}

	ctx.IndentedJSON(http.StatusOK, blobs)
}