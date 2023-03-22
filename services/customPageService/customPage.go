package customPageService

import (
	"context"

	"github.com/kzmijak/zswod_api_go/services/customPageService/selectors"
)

type CustomPageService struct {
	selectors selectors.CustomPageSelector
	ctx context.Context
}

func New(ctx context.Context) CustomPageService {
	return CustomPageService{
		selectors: selectors.Initialize(ctx),
		ctx: ctx,
	}
}