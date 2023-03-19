package imageModel

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/pickers/imageEntityPicker"
	arraymap "github.com/kzmijak/zswod_api_go/utils/arrayMap"
)

type Image struct {
	ID         uuid.UUID `json:"id,omitempty"`
	CreateTime time.Time `json:"createTime"`
	Alt        string    `json:"alt,omitempty"`
	Order 		 int      `json:"order,omitempty"`
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
		Alt: imageEntity.Alt,
		Order: imageEntity.Order,
		CreateTime: imageEntity.CreateTime,
		BlobId: blobId,
	}, nil
}

var ArrayFromEntities = arraymap.CreateArrayMapper(FromEntity)