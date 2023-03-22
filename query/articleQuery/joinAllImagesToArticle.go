package articleQuery

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/query/imageQuery"
)

func (query ArticleQuery) JoinAllImagesToArticle() ArticleQuery {
	query.WithGallery(func(gq *ent.GalleryQuery) {
		gq.WithImages(func(iq *ent.ImageQuery) {
		imageQuery.FromQuery(iq).QueryOrderedWithBlobId()
	})})

	return query
}