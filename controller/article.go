package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/services/article"
)

const (
	ErrInvalidOffset = "ErrInvalidOffset: Offset is invalid"
	ErrInvalidAmount = "ErrInvalidAmount: Amount is invalid"
)


func (c *Controller) CreateArticle(ctx *gin.Context) {
	var requestBody article.CreateArticleRequest
	var err error
	var status = http.StatusBadRequest
	defer utils.HandleError(&err, ctx, &status)

	tx, _ := database.Client.Tx(c.ctx)
	defer tx.Rollback()

	if err = ctx.BindJSON(&requestBody); err != nil {
		return 
	}

	article, err := c.articleService.CreateArticle(requestBody, tx)

	if err != nil {
		return
	}

	tx.Commit()
	ctx.IndentedJSON(http.StatusOK, article.ID)
}

func (c *Controller) UpdateArticleByTitle(ctx *gin.Context) {
	var requestBody article.UpdateArticleRequest
	var err error
	var status = http.StatusBadRequest
	defer utils.HandleError(&err, ctx, &status)

	tx, _ := database.Client.Tx(c.ctx)
	defer tx.Rollback()

	if err = ctx.BindJSON(&requestBody); err != nil {
		return
	}

	article, err := c.articleService.UpdateArticle(requestBody, tx)

	if err != nil {
		return
	}

	tx.Commit()
	ctx.IndentedJSON(http.StatusOK, article.ID)
}

func (c *Controller) GetArticleByTitle(ctx *gin.Context) {
	titleString := ctx.Param("title")

	tx, _ := database.Client.Tx(c.ctx)
	defer tx.Rollback()

	response, err := c.articleService.GetArticleByTitle(titleString, tx)

	if err != nil {
		utils.HandleError(&err, ctx)
		return
	}

	tx.Commit()
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
