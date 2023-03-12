package blob

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/blob"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrBlobIdNotFound = "ErrBlobIdNotFound: Could not find Blob by given ID"
)

func (s BlobService) ValidateBlobExistence(blobId uuid.UUID, tx *ent.Tx) error {
	exists, err := tx.Blob.Query().Where(blob.ID(blobId)).Exist(s.ctx)
	if err != nil || !exists {
		return errors.Error(ErrBlobIdNotFound)
	}

	return nil
}