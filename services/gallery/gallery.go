package gallery

import "context"

type GalleryService struct {
	ctx context.Context
}

func New() GalleryService {
	return GalleryService{}
}

func (s GalleryService) WithContext(ctx context.Context) GalleryService {
	s.ctx = ctx
	return s
}
