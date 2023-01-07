package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
}
type CreateArticleBody struct {
	Article article.CreateArticleRequest `json:"article"`
	Images []ImageBody `json:"images"`
}

func (c *Controller) CreateArticle(ctx *gin.Context) {
	var requestBody CreateArticleBody
	var err error
	var status = http.StatusBadRequest
	defer handleError(&err, ctx, &status)

	tx, err := database.Client.Tx(c.ctx)
	defer tx.Rollback()

	if err != nil {
		status = http.StatusConflict
		return
	}

	if err := ctx.BindJSON(&requestBody); err != nil {
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
		handleError(&err, ctx)
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
}

type GetArticleHeadersRequest struct {
	Offset int `form:"offset"`
	Amount int `form:"amount"`
}
func (c *Controller) GetArticleHeadersList(ctx *gin.Context) {
	var query GetArticleHeadersRequest
	var err error
	defer handleError(&err, ctx)
	
	if err = ctx.ShouldBindQuery(&query); err != nil {
		return
	}
	
	headers, err := c.articleService.GetArticleHeaders(query.Amount, query.Offset)
	if err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusOK, headers)
}