package galleryHeaderModel

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/imageModel"
	"github.com/kzmijak/zswod_api_go/pickers/galleryEntityPicker"
	"github.com/kzmijak/zswod_api_go/utils/array"
	arraymap "github.com/kzmijak/zswod_api_go/utils/arrayMap"
)

type GalleryHeaderModel struct {
	ID                   uuid.UUID          `json:"id"`
	Title                string             `json:"title"`
	CreateTime           time.Time          `json:"createTime"`
	PreviewImages		     []imageModel.Image `json:"previewImages"`
	RemainingImagesCount int								`json:"remainingImagesCount,omitempty"`
}


var Nil = GalleryHeaderModel{}

func FromGalleryEntity(galleryEntity *ent.Gallery) (GalleryHeaderModel, error) {
	galleryImageEntities, err := galleryEntityPicker.PickImageEntities(galleryEntity)
	if err != nil {
		return Nil, err
	}

	previewImageEntities := extractPreviewImages(galleryImageEntities) 
	remainingImagesCount := calculateRemainingImagesCount(len(galleryImageEntities))

	previewImageModels, err := imageModel.ArrayFromEntities(previewImageEntities)
	if err != nil {
		return Nil, err
	}

	return GalleryHeaderModel{
		ID: galleryEntity.ID,
		Title: galleryEntity.Title,
		CreateTime: galleryEntity.CreateTime,
		PreviewImages: previewImageModels,
		RemainingImagesCount: remainingImagesCount,
	}, nil
}


const MAX_IMAGES_COUNT = 6
func extractPreviewImages(imageEntities []*ent.Image) []*ent.Image {
	return array.Slice(imageEntities, 0, MAX_IMAGES_COUNT)
}

func calculateRemainingImagesCount(totalImagesCount int) int {
	if (totalImagesCount <= MAX_IMAGES_COUNT) {
		return 0
	}
	return totalImagesCount - MAX_IMAGES_COUNT
}
var ArrayFromGalleryEntities = arraymap.CreateArrayMapper(FromGalleryEntity)