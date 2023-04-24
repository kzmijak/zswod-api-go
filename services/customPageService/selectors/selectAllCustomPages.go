package selectors

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/custompage"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotQueryCustomPages = "ErrCouldNotQueryCustomPages: Failed to query for custom pages"
)

func (s CustomPageSelector) SelectAllCustomPages(tx *ent.Tx) ([]*ent.CustomPage, error) {
	customPageEntities, err := tx.CustomPage.Query().Order(ent.Asc(custompage.FieldCreateTime)).All(s.ctx)
	if err != nil {
		return nil, errors.Error(ErrCouldNotQueryCustomPages)
	}

	return customPageEntities, nil
}