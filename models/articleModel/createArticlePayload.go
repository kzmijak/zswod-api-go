package articleModel

import "github.com/google/uuid"

type CreateArticlePayload struct {
	Title     string    `json:"title"`
	Short     string    `json:"short"`
	Content   string    `json:"content"`
	GalleryId uuid.UUID `json:"galleryId"`
}

func NewCreateArticlePayload () CreateArticlePayload {
	return CreateArticlePayload{}
}

func (model CreateArticlePayload) WithTitle(title string) CreateArticlePayload {
	model.Title = title
	return model
} 

func (model CreateArticlePayload) WithShort(short string) CreateArticlePayload {
	model.Short = short
	return model
} 

func (model CreateArticlePayload) WithContent(content string) CreateArticlePayload {
	model.Content = content
	return model
} 

func (model CreateArticlePayload) WithGalleryId(galleryId uuid.UUID) CreateArticlePayload {
	model.GalleryId = galleryId
	return model
} 
