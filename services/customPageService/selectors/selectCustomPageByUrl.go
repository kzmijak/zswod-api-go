package selectors

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/customPageModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/query/customPageQuery"
)

const (
	ErrCouldNotQueryCustomPagesByUrl = "ErrCouldNotQueryCustomPagesByUrl: Could not find custom page by url" 
)

func (s CustomPageSelector) SelectCustomPageByUrl(tx *ent.Tx, url string) (model customPageModel.CustomPageModel, err error) {
	entity, err := customPageQuery.FromTx(tx).QueryByUrl(url).Only(s.ctx)
	if err != nil {
		return customPageModel.Nil, errors.Error(ErrCouldNotQueryCustomPagesByUrl)
	}

	return customPageModel.FromCustomPageEntity(entity)
}