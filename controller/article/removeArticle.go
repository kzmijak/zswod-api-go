package articleController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/utils/parser"
)

func (c ArticleController) RemoveArticle(ctx *gin.Context) {
	articleId := ctx.Param("articleId")
	var err error
	defer utils.HandleError(&err, ctx)
	
	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()

	articleIdParsed, err := parser.ParseUuid(articleId)
	if err != nil {
		return
	}
	
	err = c.ArticleService.DeleteArticle(articleIdParsed, tx)
	if err != nil {
		return
	}

	tx.Commit();
	ctx.JSON(http.StatusOK, articleIdParsed)
}