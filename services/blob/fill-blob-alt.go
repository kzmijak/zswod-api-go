package blob

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/blob"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotUpdate = "ErrCouldNotUpdate: Failed to update blob"
)

func (s BlobService) FillBlobAlt(blobId uuid.UUID, alt string, tx *ent.Tx) error {
	if err := s.ValidateBlobExistence(blobId, tx); err != nil {
		return  err
	}

	err := tx.Blob.Update().Where(blob.ID(blobId), blob.Alt("")).SetAlt(alt).Exec(s.ctx)
	if err != nil {
		return  errors.Error(ErrCouldNotUpdate)
	}

	return  nil
}