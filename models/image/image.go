package image

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	arraymap "github.com/kzmijak/zswod_api_go/utils/arrayMap"
)

type Image struct {
	ID         uuid.UUID `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Alt        string    `json:"alt,omitempty"`
	IsPreview  bool       `json:"isPreview,omitempty"`
	BlobId		 uuid.UUID `json:"blobId"`
}

func FromEntity(e *ent.Image) Image {
	return Image{
		ID: e.ID,
		Title: e.Title,
		Alt: e.Alt,
		IsPreview: e.IsPreview,
		BlobId: e.Edges.Blob.ID,
	}
}

var ArrayFromEntities = arraymap.CreateArrayMapper(FromEntity)