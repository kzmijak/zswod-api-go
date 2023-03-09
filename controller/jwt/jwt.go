package jwtController

import (
	controller "github.com/kzmijak/zswod_api_go/controller/model"
)

type JwtController struct {
	*controller.Controller
}

func New(controller *controller.Controller) *JwtController {
	return &JwtController{controller}
}
