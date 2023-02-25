package articleRepo

import (
	"context"

	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/blob"
	imageRepo "github.com/kzmijak/zswod_api_go/repositories/image"
)

type Test struct {
	ID string `json:"id"`
}

func (query *ArticleQuery) JoinAllImagesToArticle(ctx context.Context) *ArticleQuery {
	articleQuery := query.WithImages(func(iq *ent.ImageQuery) {
		imageRepo.QueryImage(iq).QueryOrderedWithBlobId(ctx).WithBlob(func(bq *ent.BlobQuery) { bq.Select(blob.FieldTitle) })
	}).Clone()

	return QueryArticle(articleQuery)
}