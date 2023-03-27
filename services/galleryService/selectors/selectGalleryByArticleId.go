package selectors

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/query/articleQuery"
)

const (
	ErrGalleryArticleIdQueryFail = "ErrGalleryArticleIdQueryFail: Could not find gallery by article ID"
)

func (s GallerySelector) SelectGalleryByArticleId(tx *ent.Tx, articleId uuid.UUID) (*ent.Gallery, error) {
	galleryId, err := articleQuery.FromTx(tx).
		QueryById(articleId).
		QueryGallery().
		OnlyID(s.ctx)
		
	if err != nil {
		return nil, errors.Error(ErrGalleryArticleIdQueryFail)
	}

	return s.GetGalleryById(tx, galleryId)
}