package galleryModel

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/imageModel"
	"github.com/kzmijak/zswod_api_go/pickers/galleryEntityPicker"
)

type GalleryModel struct {
	ID        uuid.UUID `json:"id"`
	CreatedTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
	Title     string `json:"title"`
	Images    []imageModel.Image `json:"images"`
}

var Nil = GalleryModel{}

func FromEntity(galleryEntity *ent.Gallery) (GalleryModel, error) {
	imageEntities, err := galleryEntityPicker.PickImageEntities(galleryEntity)
	if err != nil {
		return Nil, err
	}

	imageModes, err := imageModel.ArrayFromEntities(imageEntities)
	if err != nil {
		return Nil, err
	}

	return GalleryModel{
		ID: galleryEntity.ID,
		Title: galleryEntity.Title,
		CreatedTime: galleryEntity.UpdateTime,
		UpdateTime: galleryEntity.CreateTime,
		Images: imageModes,
	}, nil
} 