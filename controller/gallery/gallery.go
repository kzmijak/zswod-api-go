package galleryController

import controller "github.com/kzmijak/zswod_api_go/controller/model"

type GalleryController struct {
	controller.Controller
}

func New(controller controller.Controller) GalleryController {
	return GalleryController{controller}
}