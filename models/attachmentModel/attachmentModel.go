package attachmentModel

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/pickers/attachmentEntityPicker"
	arraymap "github.com/kzmijak/zswod_api_go/utils/arrayMap"
)

type AttachmentModel struct {
	ID          uuid.UUID `json:"id"`
	IconId      string    `json:"icon"`
	Title       string    `json:"title"`
	Description string    `json:"attachment"`
	BlobId      uuid.UUID    `json:"blobId"`
}

var Nil = AttachmentModel{}

func FromEntity(attachmentEntity *ent.Attachment) (AttachmentModel, error) {
	blob, err := attachmentEntityPicker.PickBlobEntity(attachmentEntity)
	if err != nil {
		return Nil, err
	}

	return AttachmentModel{
		ID: attachmentEntity.ID,
		IconId: attachmentEntity.IconId,
		Title: attachmentEntity.Title,
		Description: attachmentEntity.Description,
		BlobId: blob.ID,
	}, nil
}

var ArrayFromEntities = arraymap.CreateArrayMapper(FromEntity)