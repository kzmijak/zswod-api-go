package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/services/article"
)

func (c *Controller) CreateArticle(ctx *gin.Context) {
	var requestBody article.CreateArticleRequest

	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := c.articleService.CreateArticle(requestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, response)
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