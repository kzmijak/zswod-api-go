package articleRepo

import (
	"context"

	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/image"
)

func (query *ArticleQuery) JoinAllImagesToArticle(ctx context.Context) *ArticleQuery {
	articleQuery := query.WithImages(func(iq *ent.ImageQuery) {
		iq.Order(ent.Desc(image.FieldIsPreview)).WithBlob().AllX(ctx)
	}).Clone()

	return QueryArticle(articleQuery)
}