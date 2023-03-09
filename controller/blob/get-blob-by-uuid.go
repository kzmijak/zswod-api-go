package blobController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (c *BlobController) GetBlobByUuid(ctx *gin.Context) {
	uuidString := ctx.Param("uuid")

	uuidParsed, err := uuid.Parse(uuidString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	blob, err := c.BlobService.GetBlob(uuidParsed)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.Data(http.StatusOK, blob.ContentType, blob.Blob)
}