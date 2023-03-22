package customPageController

import controller "github.com/kzmijak/zswod_api_go/controller/model"

type CustomPageController struct {
	controller.Controller
}

func New(controller controller.Controller) CustomPageController {
	return CustomPageController{controller}
}