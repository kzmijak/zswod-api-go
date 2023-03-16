package imageModel

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/pickers/imageEntityPicker"
	arraymap "github.com/kzmijak/zswod_api_go/utils/arrayMap"
)

type Image struct {
	ID         uuid.UUID `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Alt        string    `json:"alt,omitempty"`
	IsPreview  bool       `json:"isPreview,omitempty"`
	BlobId		 uuid.UUID `json:"blobId"`
}

var Nil = Image{}

func FromEntity(imageEntity *ent.Image) (Image, error) {
	blobId, err := imageEntityPicker.PickImageBlobId(imageEntity)
	if err != nil {
		return Nil, err
	}

	return Image{
		ID: imageEntity.ID,
		Title: imageEntity.Title,
		Alt: imageEntity.Alt,
		IsPreview: imageEntity.IsPreview,
		BlobId: blobId,
	}, nil
}

var ArrayFromEntities = arraymap.CreateArrayMapper(FromEntity)