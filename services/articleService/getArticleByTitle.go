package articleService

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/articleModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrArticleNotFound = "ErrArticleNotFound: Could not find article"
)

func (s ArticleService) GetArticleByTitle(titleNormalized string, tx *ent.Tx) (articleModel.ArticleModel, error) {
	articleEntity, err := s.selectors.SelectFullArticleByTitleNormalized(tx, titleNormalized)
	if err != nil {
		return articleModel.Nil, errors.Error(ErrArticleNotFound)
	}

	return articleModel.FromEntity(articleEntity)
}