package gallery

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/gallery"
	"github.com/kzmijak/zswod_api_go/models/galleryHeaderModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	galleryRepo "github.com/kzmijak/zswod_api_go/repositories/gallery"
)

const (
	ErrGalleryHeadersQueryFail = "ErrGalleryHeadersQueryFail: Failed to execute gallery query"
)


func (s GalleryService) GetGalleryHeaders(amount int, offset int, tx *ent.Tx) ([]galleryHeaderModel.GalleryHeaderModel, error) {
	galleryEntities, err := galleryRepo.GalleryTx(tx).
		JoinPreviewImage().
		Order(ent.Desc(gallery.FieldCreatedAt)).
		Limit(amount).
		Offset(offset).
		All(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrGalleryHeadersQueryFail)
	}

	return galleryHeaderModel.ArrayFromGalleryEntities(galleryEntities)
}