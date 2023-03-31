package selectors

import "context"

type BlobSelector struct {
	ctx context.Context
}

func Initialize (ctx context.Context) BlobSelector {
	return BlobSelector{
		ctx: ctx,
	}
}