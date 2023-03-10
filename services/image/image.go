package image

import (
	"context"

	"github.com/kzmijak/zswod_api_go/services/blob"
)

type ImageService struct {
	ctx context.Context
	blobService *blob.BlobService
}

func New() ImageService {
	return ImageService{}
} 

func (s ImageService) WithContext(ctx context.Context) ImageService {
	s.ctx = ctx
	return s
}

func (s ImageService) WithBlobService(blobService *blob.BlobService) ImageService {
	s.blobService = blobService
	return s
}