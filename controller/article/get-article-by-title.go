package articleController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/database"
)

func (c *ArticleController) GetArticleByTitle(ctx *gin.Context) {
	titleString := ctx.Param("title")

	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()

	response, err := c.ArticleService.GetArticleByTitle(titleString, tx)

	if err != nil {
		utils.HandleError(&err, ctx)
		return
	}

	tx.Commit()
	ctx.IndentedJSON(http.StatusOK, response)
}