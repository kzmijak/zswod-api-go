package article

import (
	"context"

	"github.com/kzmijak/zswod_api_go/services/image"
)

type ArticleService struct {
	ctx context.Context
	imageService *image.ImageService
}

func New() *ArticleService {
	return &ArticleService{}
}

func (s *ArticleService) WithImageService(imageService *image.ImageService) (*ArticleService) {
	s.imageService = imageService
	return s
} 

func (s *ArticleService) WithContext(ctx context.Context) (*ArticleService) {
	s.ctx = ctx
	return s
}