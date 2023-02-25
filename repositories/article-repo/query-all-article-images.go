package articleRepo

import (
	"context"

	"github.com/kzmijak/zswod_api_go/ent"
	imageRepo "github.com/kzmijak/zswod_api_go/repositories/image-repo"
)

func (query *ArticleQuery) JoinAllImagesToArticle(ctx context.Context) *ArticleQuery {
	articleQuery := query.WithImages(func(iq *ent.ImageQuery) {
		imageRepo.QueryImage(iq).QueryOrderedWithBlobId(ctx).AllX(ctx)
	}).Clone()

	return QueryArticle(articleQuery)
}