package image

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrFailedCreatingImage = "ErrFailedCreatingImage: Failed to create image"
)

type CreateImagePayload struct {
	Title     string    `json:"title"`
	Alt       string    `json:"alt"`
	BlobId    uuid.UUID `json:"blobId"`
	IsPreview bool      `json:"isPreview"`
}

func (s ImageService) CreateImage(req CreateImagePayload, galleryId uuid.UUID, tx *ent.Tx) (*ent.Image, error) {
	image, err := tx.Image.Create().
		SetID(uuid.New()).
		SetAlt(req.Alt).
		SetGalleryID(galleryId).
		SetBlobID(req.BlobId).
		Save(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrFailedCreatingImage)
	}

	return image, nil
}