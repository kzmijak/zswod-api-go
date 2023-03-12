package blob

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrBlobDoesNotExist = "ErrBlobDoesNotExist: Requested blob does not exist"
)

func (s BlobService) GetBlob(uuid uuid.UUID, tx *ent.Tx) (*ent.Blob, error) {
	blob, err := tx.Blob.Get(s.ctx, uuid)
	if err != nil {
		return nil, errors.Error(ErrBlobDoesNotExist)
	}

	return blob, nil
}