package controller

import (
	"net/http"
	"time"

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

type BlobResponse struct {
	Title string `json:"title"`
	Alt string `json:"alt"`
	Id string `json:"id"`
	CreatedAt time.Time `json:"createdAt"` 
}

func (c *Controller) UploadBlob(ctx *gin.Context) {
	c.log.Trace("Uploading blob")

	form, err := ctx.MultipartForm()

	files := form.File["files[]"]
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrInvalidFile)
		return
	}

	var blobs []BlobResponse

	for _, file := range files {
		response, err := c.blobService.StoreBlob(file)
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
	Blobs []BlobResponse `json:"blobs"`
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

	blobUrls := lo.Map(response.Blobs, func(item *ent.Blob, index int) BlobResponse {
		return BlobResponse{
			Title: item.Title,
			Id: item.ID.String(),
			Alt: item.Alt,
			CreatedAt: item.CreatedAt,
		}
	})

	ctx.IndentedJSON(http.StatusOK, GetBlobsListResponse {
		Blobs: blobUrls,
		Eof: response.Eof,
	})
}

