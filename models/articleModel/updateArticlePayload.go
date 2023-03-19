package articleModel

type UpdateArticlePayload struct {
	Title     string    `json:"title"`
	Short     string    `json:"short"`
	Content   string    `json:"content"`
}