package imageRepo

import "github.com/kzmijak/zswod_api_go/ent"

type ImageQuery struct {
	ent.ImageQuery
}

func QueryImage (query *ent.ImageQuery) *ImageQuery {
	return &ImageQuery{
		*query,
	}
}
