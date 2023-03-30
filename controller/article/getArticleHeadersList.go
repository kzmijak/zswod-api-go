package articleController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/database"
)

func (c ArticleController) GetArticleHeadersList(ctx *gin.Context) {
	var query utils.PaginationQuery
	var err error
	defer utils.HandleError(&err, ctx)
	
	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()

	if err = ctx.ShouldBindQuery(&query); err != nil {
		return
	}
	
	headers, err := c.ArticleService.GetArticleHeaders(query.Amount, query.Offset, tx)
	if err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusOK, headers)
}
