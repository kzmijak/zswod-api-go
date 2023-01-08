package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/services/article"
	"github.com/kzmijak/zswod_api_go/services/image"
)

const (
	ErrInvalidOffset = "ErrInvalidOffset: Offset is invalid"
	ErrInvalidAmount = "ErrInvalidAmount: Amount is invalid"
)

type ImageBody struct {
	Title string `json:"title"`
	Alt string `json:"alt"`		
	Url string `json:"url"`
	Order string `json:"order"`
}
type CreateArticleBody struct {
	Article article.CreateArticleRequest `json:"article"`
	Images []ImageBody `json:"images" binding:"min=1"`
}

func (c *Controller) CreateArticle(ctx *gin.Context) {
	var requestBody CreateArticleBody
	var err error
	var status = http.StatusBadRequest
	defer utils.HandleError(&err, ctx, &status)

	tx, err := database.Client.Tx(c.ctx)
	defer tx.Rollback()

	if err != nil {
		status = http.StatusConflict
		return
	}

	if err = ctx.BindJSON(&requestBody); err != nil {
		return 
	}

	article, err := c.articleService.CreateArticle(requestBody.Article)

	if err != nil {
		return
	}

	for _, img := range requestBody.Images {
		_, err = c.imageService.CreateImage(image.CreateImageRequest{
			Title: img.Title,
			Alt: img.Alt,
			Url: img.Url,
			ArticleId: article.ID,
		})

		if err != nil {
			return
		}
	}

	tx.Commit()
	ctx.IndentedJSON(http.StatusOK, article.ID)
}

func (c *Controller) GetArticleByTitle(ctx *gin.Context) {
	titleString := ctx.Param("title")

	response, err := c.articleService.GetArticleByTitle(titleString)

	if err != nil {
		utils.HandleError(&err, ctx)
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
}


func (c *Controller) GetArticleHeadersList(ctx *gin.Context) {
	var query utils.PaginationQuery
	var err error
	defer utils.HandleError(&err, ctx)
	
	if err = ctx.ShouldBindQuery(&query); err != nil {
		return
	}
	
	headers, err := c.articleService.GetArticleHeaders(query.Amount, query.Offset)
	if err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusOK, headers)
}

func (c *Controller) GetArticleGalleriesList(ctx *gin.Context) {
	var query utils.PaginationQuery
	var err error
	defer utils.HandleError(&err, ctx)

	if err = ctx.ShouldBindQuery(&query); err != nil {
		return
	}

	galleries, err := c.articleService.GetArticleGalleries(query.Amount, query.Offset)
	if err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusOK, galleries)
}