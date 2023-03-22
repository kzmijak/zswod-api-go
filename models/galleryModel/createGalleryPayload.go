package galleryModel

import "github.com/google/uuid"

type CreateGalleryPayload struct {
	Title    string `json:"title"`
	AuthorId uuid.UUID `json:"authorId"`
}

func NewCreateGalleryPayload() CreateGalleryPayload {
	return CreateGalleryPayload{}
}

func (payload CreateGalleryPayload) WithTitle(title string) CreateGalleryPayload {
	payload.Title = title
	return payload
}

func (payload CreateGalleryPayload) WithAuthorId(authorId uuid.UUID) CreateGalleryPayload {
	payload.AuthorId = authorId
	return payload
}