package articleService

import (
	"context"

	"github.com/kzmijak/zswod_api_go/services/articleService/selectors"
)

type ArticleService struct {
	selectors selectors.ArticleSelector
	ctx context.Context
}

func New(ctx context.Context) ArticleService {
	return ArticleService{
		selectors: selectors.Initialize(ctx),
		ctx: ctx,
	}
}
