package article

import (
	"context"

	"github.com/kzmijak/zswod_api_go/services/image"
)

type ArticleService struct {
	ctx context.Context
	imageService image.ImageService
}

func New(ctx context.Context) ArticleService {
	return ArticleService{
		ctx: ctx,
	}
}

func (s ArticleService) WithImageService(imageService image.ImageService) (ArticleService) {
	s.imageService = imageService
	return s
} 