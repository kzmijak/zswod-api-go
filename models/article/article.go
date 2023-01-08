package article

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
)

type Article struct {
	ID         uuid.UUID     `json:"id,omitempty"`
	Title      string        `json:"title,omitempty"`
	Short      string        `json:"short,omitempty"`
	Content    string        `json:"content,omitempty"`
	UploadDate time.Time     `json:"uploadDate,omitempty"`
}

func FromEntity(e ent.Article) Article {
	return Article{
		ID: e.ID,
		Title: e.Title,
		Short: e.Short,
		Content: e.Content,
		UploadDate: e.UploadDate,
	}
}