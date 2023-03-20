package attachmentModel

import "github.com/google/uuid"

type AttachmentModel struct {
	ID          uuid.UUID `json:"id"`
	IconId      string    `json:"icon"`
	Title       string    `json:"title"`
	Description string    `json:"attachment"`
	BlobId      string    `json:"blobId"`
}