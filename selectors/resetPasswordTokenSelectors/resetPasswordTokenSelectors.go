package resetPasswordTokenSelectors

import "context"

type ResetPasswordTokenSelectors struct {
	ctx context.Context
}

func Initialize(ctx context.Context) ResetPasswordTokenSelectors {
	return ResetPasswordTokenSelectors{
		ctx: ctx,
	}
}