package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/services/article"
	"github.com/kzmijak/zswod_api_go/services/image"
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

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	article, err := c.articleService.CreateArticle(requestBody.Article)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
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
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
	}

	ctx.IndentedJSON(http.StatusOK, article.ID)
}

func (c *Controller) GetArticleByTitle(ctx *gin.Context) {
	titleString := ctx.Param("title")

	response, err := c.articleService.GetArticleByTitle(titleString)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
}