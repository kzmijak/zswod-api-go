package selectors

import "context"

type GallerySelector struct {
	ctx context.Context
}

func Initialize(ctx context.Context) GallerySelector {
	return GallerySelector{
		ctx: ctx,
	}
}