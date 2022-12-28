package blob

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/blob"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrBlobDoesNotExist = "ErrBlobDoesntExist: Requested blob does not exist"
)

func (s BlobService) GetBlob(uuid uuid.UUID) (*ent.Blob, error) {
	blob, err := database.Client.Blob.Query().
		Where(blob.ID(uuid)).
		Only(s.ctx);

	if err != nil {
		return nil, errors.Error(ErrBlobDoesNotExist)
	}

	return blob, nil
}