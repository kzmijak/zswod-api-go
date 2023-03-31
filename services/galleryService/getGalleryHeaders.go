package galleryService

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/gallery"
	"github.com/kzmijak/zswod_api_go/models/galleryHeaderModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/query/galleryQuery"
)

const (
	ErrGalleryHeadersQueryFail = "ErrGalleryHeadersQueryFail: Failed to execute gallery query"
)


func (s GalleryService) GetGalleryHeaders(amount int, offset int, tx *ent.Tx) ([]galleryHeaderModel.GalleryHeaderModel, error) {
	galleryEntities, err := galleryQuery.FromTx(tx).
		JoinAllImagesToGallery().
		Order(ent.Desc(gallery.FieldCreateTime)).
		Limit(amount).
		Offset(offset).
		All(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrGalleryHeadersQueryFail)
	}

	return galleryHeaderModel.ArrayFromGalleryEntities(galleryEntities)
}