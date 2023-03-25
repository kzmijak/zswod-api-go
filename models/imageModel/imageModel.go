package imageModel

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	arraymap "github.com/kzmijak/zswod_api_go/utils/arrayMap"
)

type Image struct {
	ID         uuid.UUID `json:"id,omitempty"`
	CreateTime time.Time `json:"createTime"`
	Alt        string    `json:"alt,omitempty"`
	Order 		 int      `json:"order,omitempty"`
	Src		 string `json:"src"`
}

var Nil = Image{}

func FromEntity(imageEntity *ent.Image) (Image, error) {


	return Image{
		ID: imageEntity.ID,
		Alt: imageEntity.Alt,
		Order: imageEntity.Order,
		CreateTime: imageEntity.CreateTime,
		Src: imageEntity.Src,
	}, nil
}

var ArrayFromEntities = arraymap.CreateArrayMapper(FromEntity)