package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	ErrInvalidFile = "ErrInvalidFile: Uploaded file is invalid"
	ErrInvalidUuidFormat = "ErrInvalidUuidFormat: Could not parse UUID. Invalid format."
)

func (c *Controller) UploadBlob(ctx *gin.Context) {
	c.log.Trace("Uploading blob")

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrInvalidFile)
		return
	}

	response, err := c.blobService.StoreBlob(file)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) GetBlobByUuid(ctx *gin.Context) {
	uuidString := ctx.Param("uuid")

	uuidParsed, err := uuid.Parse(uuidString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	blob, err := c.blobService.GetBlob(uuidParsed)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.Data(http.StatusOK, blob.ContentType, blob.Blob)

}