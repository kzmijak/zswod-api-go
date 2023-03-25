package blobModel

import (
	"time"

	"github.com/google/uuid"
)

type BlobModel struct {
	ID          uuid.UUID `json:"id"`
	CreateTime  time.Time `json:"createTime"`
	Blob        []byte    `json:"blob"`
	ContentType string    `json:"contentType"`
}