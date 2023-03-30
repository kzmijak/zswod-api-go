package articleController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/models/articleModel"
	"github.com/kzmijak/zswod_api_go/models/imageModel"
	"github.com/kzmijak/zswod_api_go/modules/database"
)

type UpdateArticleRequest struct {
	TitleNormalized string `json:"titleNormalized"`
	Article articleModel.UpdateArticlePayload `json:"article"`
	Images []imageModel.CreateImagePayload `json:"images"`
}

func (c ArticleController) UpdateArticle(ctx *gin.Context) {
	var requestBody UpdateArticleRequest
	var err error
	var status = http.StatusBadRequest
	defer utils.HandleError(&err, ctx, &status)
	
	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()
	
	if err = ctx.BindJSON(&requestBody); err != nil {
		return
	}

	article, err := c.ArticleService.GetArticleByTitle(requestBody.TitleNormalized, tx)
	if err != nil {
		return
	}

	articleUpdated, err := c.ArticleService.UpdateArticle(article.ID, requestBody.Article, tx)
	if err != nil {
		return
	}

	relatedGallery, err := c.GalleryService.GetGalleryByArticleId(articleUpdated.ID, tx)
	if err != nil {
		return
	}

	err = c.GalleryService.FlushGalleryImages(relatedGallery.ID, tx)
	if err != nil {
		return 
	}
	
	for _, imageRequest := range requestBody.Images {
		_, err := c.ImageService.CreateImage(imageRequest, relatedGallery.ID, tx)
		if err != nil {
			return
		}
	}

	tx.Commit()
	ctx.IndentedJSON(http.StatusOK, articleUpdated.ID)
}