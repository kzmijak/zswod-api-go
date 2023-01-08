package image

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	arraymap "github.com/kzmijak/zswod_api_go/utils/arrayMap"
)

type Image struct {
	ID         uuid.UUID `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Alt        string    `json:"alt,omitempty"`
	URL        string    `json:"url,omitempty"`
	UploadDate time.Time `json:"uploadDate,omitempty"`
	Order      int       `json:"order,omitempty"`
}

func FromEntity(e *ent.Image) Image {
	return Image{
		ID: e.ID,
		Title: e.Title,
		Alt: e.Alt,
		URL: e.URL,
		UploadDate: e.UploadDate,
		Order: e.Order,
	}
}

var ArrayFromEntities = arraymap.CreateArrayMapper(FromEntity)