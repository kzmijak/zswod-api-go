package blob

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotUpdate = "ErrCouldNotUpdate: Failed to update blob"
)

func (s BlobService) UpdateBlob(target *ent.Blob) (*ent.Blob, error) {
	blob, err := database.Client.Blob.UpdateOne(target).Save(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrCouldNotUpdate)
	}

	return blob, nil
}