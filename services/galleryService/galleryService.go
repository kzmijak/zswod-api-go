package galleryService

import (
	"context"

	"github.com/kzmijak/zswod_api_go/services/galleryService/selectors"
)

type GalleryService struct {
	selectors selectors.GallerySelector
	ctx context.Context
}

func New(ctx context.Context) GalleryService {
	return GalleryService{
		selectors: selectors.Initialize(ctx),
		ctx: ctx,
	}
}

