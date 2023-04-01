package blobModel

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/blobType"
	arraymap "github.com/kzmijak/zswod_api_go/utils/arrayMap"
)

type BlobModel struct {
	ID          uuid.UUID  `json:"id"`
	Title				string		 `json:"title"`
	CreateTime  time.Time  `json:"createTime"`
	ContentType string     `json:"contentType"`
	IsPublic 	  bool 			 `json:"isPublic"`
	Type blobType.BlobType `json:"blobType"`
	blob        []byte
}

var Nil = BlobModel{}

func FromEntity(blobEntity *ent.Blob) (BlobModel, error) {
	blobType, err := blobType.FromString(blobEntity.Type.String())
	if err != nil {
		return Nil, err
	}

	return BlobModel{
		ID: blobEntity.ID,
		CreateTime: blobEntity.CreateTime,
		ContentType: blobEntity.ContentType,
		IsPublic: blobEntity.IsPublic,
		Type: blobType,
		blob: blobEntity.Blob,
		Title: blobEntity.Title,
	}, nil
}

func (blobModel BlobModel) GetBytes() []byte {
	return blobModel.blob
}

var ArrayFromEntities = arraymap.CreateArrayMapper(FromEntity)
