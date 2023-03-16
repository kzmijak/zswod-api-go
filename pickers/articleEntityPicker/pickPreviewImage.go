package articleEntityPicker

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/pickers/galleryEntityPicker"
)

func PickGalleryPreviewImage(articleEntity *ent.Article) (*ent.Image, error) {
	galleryEntity, err := PickGalleryEntity(articleEntity)
	if err != nil {
		return nil, err
	}

	previewImageEntity, err := galleryEntityPicker.PickPreviewImage(galleryEntity)
	if err != nil {
		return nil, err
	}

	return previewImageEntity, nil
}