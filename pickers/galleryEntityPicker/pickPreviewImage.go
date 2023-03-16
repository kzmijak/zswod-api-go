package galleryEntityPicker

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrGalleryPreviewImageNotFound = "ErrGalleryPreviewImageNotFound: Could not find preview image for specified gallery"
)

func PickPreviewImage(gallery *ent.Gallery) (*ent.Image, error) {
	imageEntities, err := PickImageEntities(gallery)
	if err != nil {
		return nil, err
	}

	var previewImageEntity *ent.Image
	for _, imageEntity := range imageEntities {
		if imageEntity.IsPreview {
			previewImageEntity = imageEntity
		}
	}

	if previewImageEntity == nil {
		return nil, errors.Error(ErrGalleryPreviewImageNotFound)
	}

	return previewImageEntity, nil
}