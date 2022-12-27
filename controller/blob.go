package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ErrInvalidFile = "ErrInvalidFile: Uploaded file is invalid"
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