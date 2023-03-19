package articleController

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/articleModel"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/services/image"
)

type CreateArticleRequest struct {
	Title     string    `json:"title"`
	Short     string    `json:"short"`
	Content   string    `json:"content"`
	Images []image.CreateImagePayload `json:"images"`
}

func (c ArticleController) CreateArticle(ctx *gin.Context) {
	var requestBody CreateArticleRequest
	var err error
	defer utils.HandleError(&err, ctx)

	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()

	if err = ctx.BindJSON(&requestBody); err != nil {
		return
	}

	createArticlePayload := articleModel.NewCreateArticlePayload().
		WithTitle(requestBody.Title).
		WithShort(requestBody.Short).
		WithContent(requestBody.Content)

	galleryId, err := c.createGallery(requestBody.Title, tx)
	if err != nil {
		return
	}

	err = c.createImages(requestBody.Images, galleryId, tx)
	if err != nil {
		return
	}

	createArticlePayload = createArticlePayload.WithGalleryId(galleryId)
	articleId, err := c.createArticle(createArticlePayload, galleryId, tx)
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

func (c ArticleController) createArticle(req articleModel.CreateArticlePayload, galleryId uuid.UUID, tx *ent.Tx) (uuid.UUID, error) {
	articleId, err := c.ArticleService.CreateArticle(req, tx)
	if err != nil {
		return uuid.Nil, err
	}
	return articleId, nil
}