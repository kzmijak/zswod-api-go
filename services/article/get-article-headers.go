package article

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/article"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrQueryFail = "ErrQueryFail: Failed to execute query"
)

type GetArticleHeadersParams struct {
	Offset int `json:"offest"`
	Amount int `json:"amount"`
}

type ArticleHeader struct {
	
}

func (s ArticleService) GetArticleHeaders(params GetArticleHeadersParams) ([]*ent.Article, error) {
	result, err := database.Client.Article.Query().
		Order(ent.Desc(article.FieldUploadDate)).
		Limit(params.Amount).
		Offset(params.Offset).
		WithTitleNormalized().
		All(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrQueryFail)
	}

	return result, nil
}