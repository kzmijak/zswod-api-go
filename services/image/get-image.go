package image

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotParseImageId = "ErrCouldNotParseImageId: Could not parse specified string to UUID"
	ErrImageIdNotFound      = "ErrImageIdNotFound: Could not find image by given ID"
)

func (s ImageService) GetImage(imageId string, tx *ent.Tx) (*ent.Image, error) {
	imageIdParsed, err := uuid.Parse(imageId)
	if err != nil {
		return nil, errors.Error(ErrCouldNotParseImageId)
	}

	image, err := tx.Image.Get(s.ctx, imageIdParsed)
	if err != nil {
		return nil, errors.Error(ErrImageIdNotFound)
	}

	return image, nil
}