package customPageService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/customPageModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotUpdateCustomPage = "ErrCouldNotUpdateCustomPage: Failed to update custom page"
)

func (s CustomPageService) UpdateCustomPage(customPageId uuid.UUID, payload customPageModel.UpdateCustomPagePayload, tx *ent.Tx) (customPageModel.CustomPageModel, error) {
	updatedEntity, err := tx.CustomPage.
		UpdateOneID(customPageId).
		SetContent(payload.Content).
		Save(s.ctx)

	if err != nil {
		return customPageModel.Nil, errors.Error(ErrCouldNotUpdateCustomPage)
	}

	return s.selectors.SelectCustomPageById(tx, updatedEntity.ID)
}