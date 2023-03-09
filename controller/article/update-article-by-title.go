package articleController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/services/article"
)

func (c *ArticleController) UpdateArticleByTitle(ctx *gin.Context) {
	var requestBody article.UpdateArticleRequest
	var err error
	var status = http.StatusBadRequest
	defer utils.HandleError(&err, ctx, &status)

	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()

	if err = ctx.BindJSON(&requestBody); err != nil {
		return
	}

	article, err := c.ArticleService.UpdateArticle(requestBody, tx)

	if err != nil {
		return
	}

	tx.Commit()
	ctx.IndentedJSON(http.StatusOK, article.ID)
}