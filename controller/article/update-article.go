package articleController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/models/articleModel"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/services/image"
	"github.com/kzmijak/zswod_api_go/utils/parser"
)

const (
	ErrCouldNotParseArticleId = "ErrCouldNotParseArticleId: Could not parse Article Guid"
)

type UpdateArticleRequest struct {
	ArticleId string `json:"articleId"`
	GalleryId string `json:"galleryId"`
	Article articleModel.UpdateArticlePayload `json:"article"`
	Images []image.CreateImagePayload `json:"images"`
}

func (c *ArticleController) UpdateArticle(ctx *gin.Context) {
	var requestBody UpdateArticleRequest
	var err error
	var status = http.StatusBadRequest
	defer utils.HandleError(&err, ctx, &status)
	
	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()
	
	if err = ctx.BindJSON(&requestBody); err != nil {
		return
	}

	articleUuid, err := parser.ParseUuid(requestBody.ArticleId)
	if err != nil {
		return
	}

	article, err := c.ArticleService.UpdateArticle(articleUuid, requestBody.Article, tx)
	if err != nil {
		return
	}

	galleryId := requestBody.GalleryId
	gallery, err := c.GalleryService.FlushGalleryImages(galleryId, tx)
	if err != nil {
		return 
	}

	
	for _, imageRequest := range requestBody.Images {
		err = c.BlobService.FillBlobAlt(imageRequest.BlobId, imageRequest.Alt, tx)
		if err != nil {
			return
		}

		_, err := c.ImageService.CreateImage(imageRequest, gallery.ID, tx)
		if err != nil {
			return
		}
	}

	tx.Commit()
	ctx.IndentedJSON(http.StatusOK, article.ID)
}