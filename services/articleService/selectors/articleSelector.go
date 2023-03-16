package selectors

import (
	"context"
)

type ArticleSelector struct {
	ctx context.Context
}

func Initialize (ctx context.Context) ArticleSelector {
	return ArticleSelector{
		ctx: ctx,
	}
}
