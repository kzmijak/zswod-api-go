package article

import (
	"context"
)

type ArticleService struct {
	ctx context.Context
}

func New(ctx context.Context) ArticleService {
	return ArticleService{
		ctx: ctx,
	}
}
