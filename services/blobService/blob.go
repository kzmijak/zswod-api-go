package blobService

import (
	"context"

	"github.com/kzmijak/zswod_api_go/services/blobService/selectors"
)

type BlobService struct {
	selectors selectors.BlobSelector
	ctx context.Context
}

func New(ctx context.Context) BlobService {
	return BlobService{
		selectors: selectors.Initialize(ctx),
		ctx: ctx,
	}
}
