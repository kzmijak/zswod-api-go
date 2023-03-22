package selectors

import "context"

type CustomPageSelector struct {
	ctx context.Context
}

func Initialize (ctx context.Context) CustomPageSelector {
	return CustomPageSelector{
		ctx: ctx,
	}
}