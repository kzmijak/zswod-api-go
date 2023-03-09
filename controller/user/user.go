package userController

import (
	controller "github.com/kzmijak/zswod_api_go/controller/model"
)

type UserController struct {
	*controller.Controller
}

func New(controller *controller.Controller) UserController {
	return UserController{controller}
}
