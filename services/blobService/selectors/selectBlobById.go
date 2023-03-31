package selectors

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/query/blobQuery"
)

const (
	ErrBlobIdNotFound = "ErrBlobIdNotFound: Could not find blob by given ID"
)


func (s BlobSelector) SelectBlobById(tx *ent.Tx, blobId uuid.UUID) (*ent.Blob, error) {
	blobEntity, err := blobQuery.FromTx(tx).
		QueryById(blobId).
		Only(s.ctx)
	
	if err != nil {
		return nil, errors.Error(ErrBlobIdNotFound)
	}

	return blobEntity, nil
}