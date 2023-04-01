package userSelectors

import "context"

type UserSelectors struct {
	ctx context.Context
}

func Initialize(ctx context.Context) UserSelectors {
	return UserSelectors{
		ctx: ctx,
	}
}