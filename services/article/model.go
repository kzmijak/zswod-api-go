package article

import "github.com/google/uuid"

type BaseArticlePayload struct {
	Title   string `json:"title"`
	Short   string `json:"short"`
	Content string `json:"content"`
}

type BaseImagePayload struct {
	Title     string `json:"title"`
	Alt       string `json:"alt"`
	BlobId    uuid.UUID `json:"blobId"`
	IsPreview bool   `json:"isPreview"`
}