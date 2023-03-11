package authController

import "github.com/kzmijak/zswod_api_go/controller"

type AuthController struct {
	controller.Controller
}

func New(controller controller.Controller) AuthController {
	return AuthController{controller}
}