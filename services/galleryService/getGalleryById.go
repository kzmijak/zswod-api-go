package galleryService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/galleryModel"
)

func (s GalleryService) GetGalleryById(galleryId uuid.UUID, tx *ent.Tx) (gallery galleryModel.GalleryModel, err error)  {
	gallery = galleryModel.Nil
	galleryEntity, err := s.selectors.GetGalleryById(tx, galleryId)
	if err != nil {
		return
	}

	gallery, err = galleryModel.FromEntity(galleryEntity)
	return
}