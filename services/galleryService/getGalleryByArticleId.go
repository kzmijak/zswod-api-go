package galleryService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/galleryModel"
)

func (s GalleryService) GetGalleryByArticleId(articleId uuid.UUID, tx *ent.Tx) (galleryModel.GalleryModel, error) {
	galleryEntity, err := s.selectors.SelectGalleryByArticleId(tx, articleId)
	if err != nil {
		return galleryModel.Nil, err
	}

	return galleryModel.FromEntity(galleryEntity)
}