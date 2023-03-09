package image

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrFailedCreatingImage = "ErrFailedCreatingImage: Failed to create image"
)

type CreateImageRequest struct {
	Title string `json:"title"`
	Alt string `json:"alt"`
	BlobId uuid.UUID `json:"blobId"`
	GalleryId uuid.UUID `json:"galleryId"`
	IsPreview bool `json:"isPreview"`
}

func (s ImageService) CreateImage(req CreateImageRequest, tx *ent.Tx) (image *ent.Image, err error) {

	blob, err :=s.blobService.GetBlob(req.BlobId)
	if err != nil {
		return
	}

	if (blob.Alt != "") {
		blob.Alt = req.Alt
		s.blobService.UpdateBlob(blob)
	}

	image, err = tx.Image.Create().
		SetID(uuid.New()).
		SetTitle(req.Title).
		SetAlt(req.Alt).
		SetGalleryID(req.GalleryId).
		SetBlobID(req.BlobId).
		SetIsPreview(req.IsPreview).
		Save(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrFailedCreatingImage)
	}

	return image, nil
}