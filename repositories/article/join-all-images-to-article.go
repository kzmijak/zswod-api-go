package articleRepo

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/gallery"
	imageRepo "github.com/kzmijak/zswod_api_go/repositories/image"
)

func (query ArticleQuery) JoinAllImagesToArticle() ArticleQuery {
	query.WithGallery(func(gq *ent.GalleryQuery) {
		gq.Select(gallery.FieldID).WithImages(func(iq *ent.ImageQuery) {
		imageRepo.QueryImage(iq).QueryOrderedWithBlobId()
	})})

	return query
}