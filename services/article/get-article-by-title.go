package article

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	articleRepo "github.com/kzmijak/zswod_api_go/repositories/article"
)

const (
	ErrArticleNotFound = "ErrArticleNotFound: Could not find article"
)

func (s ArticleService) GetArticleByTitle(titleNormalized string, tx *ent.Tx) (*ent.Article, error) {
	articleEntity, err := articleRepo.ArticleTx(tx).
		QueryArticleEntityByTitle(titleNormalized).
		Only(s.ctx)

	if err != nil {
		return nil, errors.Error(ErrArticleNotFound)
	}

	return articleEntity, nil
}