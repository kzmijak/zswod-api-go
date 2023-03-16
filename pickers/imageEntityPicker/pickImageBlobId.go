package imageEntityPicker

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotPickBlobFromImage = "ErrCouldNotPickBlobFromImage: Failed to pick blob from image"
)

func PickImageBlobId(imageEntity *ent.Image) (uuid.UUID, error) {
	blob, err := imageEntity.Edges.BlobOrErr()
	if err != nil {
		return uuid.Nil, errors.Error(ErrCouldNotPickBlobFromImage)
	}

	return blob.ID, nil
}