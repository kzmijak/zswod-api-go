package articleRepo

import (
	"github.com/kzmijak/zswod_api_go/ent"
	imageRepo "github.com/kzmijak/zswod_api_go/repositories/image"
)

type Test struct {
	ID string `json:"id"`
}

func (query ArticleQuery) JoinAllImagesToArticle() ArticleQuery {
	query.WithGallery(func(gq *ent.GalleryQuery) {
		gq.WithImages(func(iq *ent.ImageQuery) {
		imageRepo.QueryImage(iq).QueryOrderedWithBlobId()
	})})

	return query
}