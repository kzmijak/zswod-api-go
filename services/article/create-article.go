package article

import (
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrFailedCreatingArticle = "ErrFailedCreatingArticle: Failed to create article"
	ErrFailedCreatingArticleTitle = "ErrFailedCreatingArticleTitle: Failed to create article title"
)

type CreateArticleRequest struct {
	Title string `json:"title"`
	Short string `json:"short"`
	Content string `json:"content"`
}


func (s ArticleService) CreateArticle(req CreateArticleRequest) (*ent.Article, error) {
	const MAX_LENGTH = 64	

	titleSanitized := sanitize(req.Title)
	if len(titleSanitized) > MAX_LENGTH {
		titleSanitized = titleSanitized[:MAX_LENGTH - 1]
	}
	
	article, err := database.Client.Article.Create().
		SetID(uuid.New()).
		SetTitle(req.Title).
		SetShort(req.Short).
		SetContent(req.Content).
		SetUploadDate(time.Now()).
		Save(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrFailedCreatingArticle)
	}
	
	_, err = database.Client.ArticleTitleGuid.
		Create().
		SetID(uuid.New()).
		SetTitleNormalized(titleSanitized).
		SetArticle(article).
		Save(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrFailedCreatingArticleTitle)
	}

	return article, nil
}

func sanitize(input string) string {
	reg, err := regexp.Compile("[ąćęłńóśźż]")
    if err != nil {
        panic(err)
    }
    input = reg.ReplaceAllString(input, "")

    input = strings.Replace(input, " ", "-", -1)

    return input
}