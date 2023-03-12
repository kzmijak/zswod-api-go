package image

import (
	"context"
)

type ImageService struct {
	ctx context.Context
}

func New() ImageService {
	return ImageService{}
} 

func (s ImageService) WithContext(ctx context.Context) ImageService {
	s.ctx = ctx
	return s
}
