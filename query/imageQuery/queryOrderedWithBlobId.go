package imageQuery

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/blob"
	"github.com/kzmijak/zswod_api_go/ent/image"
)

func (query ImageQuery) QueryOrderedWithBlobId() ImageQuery {
	query.Order(ent.Desc(image.FieldIsPreview)).WithBlob(func(bq *ent.BlobQuery) { bq.Select(blob.FieldID) }).Clone()

	return query
}
