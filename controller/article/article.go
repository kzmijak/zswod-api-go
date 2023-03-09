package articleController

import (
	controller "github.com/kzmijak/zswod_api_go/controller/model"
)

type ArticleController struct {
	*controller.Controller
}

func New(controller *controller.Controller) ArticleController {
	return ArticleController{controller}
}
