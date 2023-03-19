package galleryHeaderModel

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/imageModel"
	"github.com/kzmijak/zswod_api_go/pickers/galleryEntityPicker"
	arraymap "github.com/kzmijak/zswod_api_go/utils/arrayMap"
)

type GalleryHeaderModel struct {
	ID           uuid.UUID        `json:"id"`
	Title        string           `json:"title"`
	CreateTime   time.Time        `json:"createTime"`
	PreviewImage imageModel.Image `json:"previewImage"`
}

var Nil = GalleryHeaderModel{}

func FromGalleryEntity(galleryEntity *ent.Gallery) (GalleryHeaderModel, error) {
	previewImageEntity, err := galleryEntityPicker.PickPreviewImage(galleryEntity)
	if err != nil {
		return Nil, err
	}

	previewImageModel, err := imageModel.FromEntity(previewImageEntity)
	if err != nil {
		return Nil, err
	}

	return GalleryHeaderModel{
		ID: galleryEntity.ID,
		Title: galleryEntity.Title,
		CreateTime: galleryEntity.CreateTime,
		PreviewImage: previewImageModel,
	}, nil
}

var ArrayFromGalleryEntities = arraymap.CreateArrayMapper(FromGalleryEntity)