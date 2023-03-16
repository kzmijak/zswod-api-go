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
	Title     string `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
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
		CreatedAt: galleryEntity.CreatedAt,
		Images: imageModes,
	}, nil
} 