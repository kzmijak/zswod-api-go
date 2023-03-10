package blobController

import (
	controller "github.com/kzmijak/zswod_api_go/controller/model"
)

type BlobController struct {
	controller.Controller
}

func New(controller controller.Controller) BlobController {
	return BlobController{controller}
}
