package articleEntityPicker

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotPickGalleryFromArticle = "ErrCouldNotPickGalleryFromArticle: Failed to pick gallery from article"
)

func PickGalleryEntity(articleEntity *ent.Article) (*ent.Gallery, error) {
	galleryEntity, err := articleEntity.Edges.GalleryOrErr()
	if err != nil {
		return nil, errors.Error(ErrCouldNotPickGalleryFromArticle)
	}

	return galleryEntity, nil
}