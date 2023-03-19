package articleModel

import "github.com/google/uuid"

type UpdateArticlePayload struct {
	Title     string    `json:"title"`
	Short     string    `json:"short"`
	Content   string    `json:"content"`
	GalleryId uuid.UUID `json:"galleryId"`
}