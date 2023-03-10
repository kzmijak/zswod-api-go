package articleController

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/services/article"
	"github.com/kzmijak/zswod_api_go/services/image"
)

type CreateArticleRequest struct {
	Article article.CreateArticlePayload `json:"article"`
	Images []image.CreateImagePayload `json:"images"`
}

func (c ArticleController) CreateArticle(ctx *gin.Context) {
	var requestBody CreateArticleRequest
	var err error
	var status = http.StatusBadRequest
	defer utils.HandleError(&err, ctx, &status)

	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()

	if err = ctx.BindJSON(&requestBody); err != nil {
		return
	}

	galleryId, err := c.createGallery(requestBody.Article.Title, tx)
	if err != nil {
		return
	}

	err = c.createImages(requestBody.Images, galleryId, tx)
	if err != nil {
		return
	}

	articleId, err := c.createArticle(requestBody.Article, galleryId, tx)
	if err != nil {
		return
	}

	tx.Commit()
	ctx.IndentedJSON(http.StatusOK, articleId)
}

func (c ArticleController) createGallery(title string, tx *ent.Tx) (uuid.UUID, error) {
	galleryName := fmt.Sprintf("Galeria artykułu \"%v\"", title)
	galleryEntity, err := c.GalleryService.CreateGallery(galleryName, tx)
	if err != nil {
		return uuid.Nil, err
	}
	return galleryEntity.ID, nil
}

func (c ArticleController) createImages(req []image.CreateImagePayload, galleryId uuid.UUID, tx *ent.Tx) error {
	for _, img := range req {
		_, err := c.ImageService.CreateImage(img, galleryId, tx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c ArticleController) createArticle(req article.CreateArticlePayload, galleryId uuid.UUID, tx *ent.Tx) (string, error) {
	articleEntity, err := c.ArticleService.CreateArticle(req, galleryId, tx)
	if err != nil {
		return "", err
	}
	return articleEntity.Title, nil
}