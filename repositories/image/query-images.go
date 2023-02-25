package imageRepo

import (
	"context"

	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/image"
)

func (query ImageQuery) QueryOrderedWithBlobId (ctx context.Context) ImageQuery {
	query.Order(ent.Desc(image.FieldIsPreview)).WithBlob(func(bq *ent.BlobQuery) { bq.IDs(ctx) }).Clone()

	return query
}