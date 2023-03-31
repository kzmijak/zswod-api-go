package blobService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/blobModel"
)

const (
	ErrBlobDoesNotExist = "ErrBlobDoesNotExist: Requested blob does not exist"
)

func (s BlobService) GetBlobById(uuid uuid.UUID, tx *ent.Tx) (blobModel.BlobModel, error) {
	blob, err := s.selectors.SelectBlobById(tx, uuid)
	if err != nil {
		return blobModel.Nil, err
	}

	return blobModel.FromEntity(blob)
}