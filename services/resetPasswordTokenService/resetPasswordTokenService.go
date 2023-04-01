package resetPasswordTokenService

import (
	"context"

	"github.com/kzmijak/zswod_api_go/selectors/resetPasswordTokenSelectors"
	"github.com/kzmijak/zswod_api_go/selectors/userSelectors"
)

type ResetPasswordTokenService struct {
	tokenSelectors resetPasswordTokenSelectors.ResetPasswordTokenSelectors
	userSelectors userSelectors.UserSelectors
	ctx       context.Context
}

func New(ctx context.Context) ResetPasswordTokenService {
	return ResetPasswordTokenService{
		tokenSelectors: resetPasswordTokenSelectors.Initialize(ctx),
		userSelectors: userSelectors.Initialize(ctx),
		ctx: ctx,
	}
}
