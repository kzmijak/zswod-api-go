package articleQuery

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/query/galleryQuery"
)

func (query ArticleQuery) JoinAllImagesToArticle() ArticleQuery {
	query.WithGallery(func(gq *ent.GalleryQuery) {
		galleryQuery.FromQuery(gq).JoinAllImagesToGallery() 
	})

	return query
}