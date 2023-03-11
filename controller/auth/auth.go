package authController

import (
	controller "github.com/kzmijak/zswod_api_go/controller/model"
)
type AuthController struct {
	controller.Controller
}

func New(controller controller.Controller) AuthController {
	return AuthController{controller}
}