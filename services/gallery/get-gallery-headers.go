package gallery

import (
	"time"

	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/gallery"
	imageMdl "github.com/kzmijak/zswod_api_go/models/image"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	galleryRepo "github.com/kzmijak/zswod_api_go/repositories/gallery"
)

const (
	ErrGalleryHeadersQueryFail = "ErrGalleryHeadersQueryFail: Failed to execute gallery query"
)

type GalleryHeader struct {
	Title      string    `json:"title"`
	UploadDate time.Time `json:"uploadDate"`
	PreviewImage imageMdl.Image `json:"previewImage"`
}

func (s GalleryService) GetGalleryHeaders(amount int, offset int, tx *ent.Tx) ([]GalleryHeader, error) {
	galleryEntities, err := galleryRepo.GalleryTx(tx).
		JoinPreviewImage().
		Order(ent.Desc(gallery.FieldCreatedAt)).
		Limit(amount).
		Offset(offset).
		All(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrGalleryHeadersQueryFail)
	}

	headers := []GalleryHeader{}
	for _, entity := range galleryEntities {
		headers = append(headers, GalleryHeader{
			Title: entity.Title,
			UploadDate: entity.CreatedAt,
			PreviewImage: imageMdl.FromEntity(entity.Edges.Images[0]),
		})
	}

	return headers, nil
}