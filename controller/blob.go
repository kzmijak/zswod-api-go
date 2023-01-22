package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/samber/lo"
)

const (
	ErrInvalidFile = "ErrInvalidFile: Uploaded file is invalid"
	ErrInvalidUuidFormat = "ErrInvalidUuidFormat: Could not parse UUID. Invalid format."
)

type BlobUrlResponse struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Id string `json:"id"`
}

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

	url := createBlobUrl(response, ctx)

	ctx.IndentedJSON(http.StatusOK, BlobUrlResponse{ 
		Name: response.Name,
		Url: url,
		Id: response.ID.String(),
	})
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

type GetBlobsListResponse struct {
	Blobs []BlobUrlResponse `json:"blobs"`
	Eof bool `json:"eof"`
}
func (c *Controller) GetBlobsList(ctx *gin.Context) {
	var query utils.PaginationQuery
	var err error
	defer utils.HandleError(&err, ctx)
	
	if err = ctx.ShouldBindQuery(&query); err != nil {
		return
	}
	
	response, err := c.blobService.GetBlobs(query.Amount, query.Offset)
	if err != nil {
		return
	}

	blobUrls := lo.Map(response.Blobs, func(item *ent.Blob, index int) BlobUrlResponse {
		return BlobUrlResponse{
			Name: item.Name,
			Url: createBlobUrl(item, ctx),
			Id: item.ID.String(),
		}
	})

	ctx.IndentedJSON(http.StatusOK, blobUrls)
}

func createBlobUrl(blob *ent.Blob, ctx *gin.Context) string {
	scheme := "http://"
	if ctx.Request.TLS != nil {
			scheme = "https://"
	}
	blobUrl := fmt.Sprint(scheme, ctx.Request.Host, ctx.Request.URL.Path, "/", blob.ID.String()) 

	return blobUrl
}
