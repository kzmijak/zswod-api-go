package galleryController

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kzmijak/zswod_api_go/controller/utils"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/utils/parser"
)

func (c GalleryController) GetGalleryById(ctx *gin.Context) {
	galleryIdString := ctx.Param("galleryId")

	var err error
	defer utils.HandleError(&err, ctx)

	tx, _ := database.Client.Tx(c.Ctx)
	defer tx.Rollback()


	galleryId, err := parser.ParseUuid(galleryIdString)

	galleryModel, err := c.GalleryService.GetGalleryById(galleryId, tx)
	ctx.JSON(http.StatusOK, galleryModel)
}