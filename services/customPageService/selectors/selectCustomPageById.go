package selectors

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/customPageModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/query/customPageQuery"
)

const (
	ErrCouldNotQueryCustomPageById = "ErrCouldNotQueryCustomPageById: Failed to query custom page by provided ID" 
)

func (s CustomPageSelector) SelectCustomPageById(tx *ent.Tx, customPageId uuid.UUID) (customPageModel.CustomPageModel, error) {
	entity, err := customPageQuery.FromTx(tx).
		QueryById(customPageId).
		Only(s.ctx)

	if err != nil {
		return customPageModel.Nil, errors.Error(ErrCouldNotQueryCustomPageById)
	}

	return customPageModel.FromCustomPageEntity(entity)
}