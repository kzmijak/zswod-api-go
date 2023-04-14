package customPageService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotRemoveCustomPageById = "ErrCouldNotRemoveCustomPageById: Failed to remove custom page by ID"
)

func (s CustomPageService) DeleteCustomPageById(id uuid.UUID, tx *ent.Tx) error {
	err := tx.CustomPage.DeleteOneID(id).Exec(s.ctx)

	if err != nil {
		return errors.Error(ErrCouldNotRemoveCustomPageById)
	}

	return nil
}