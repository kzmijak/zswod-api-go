package blob

import "context"

type BlobService struct {
	ctx context.Context
}

func New() *BlobService {
	return &BlobService{}
}

func (s BlobService) WithContext(ctx context.Context) (*BlobService) {
	s.ctx = ctx
	return &s
}