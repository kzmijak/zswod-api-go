package selectors

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/article"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrArticleTitleNormalizedNotFound = "ErrArticleTitleNormalizedNotFound: Could not find article by given Title (normalized)"
)

func (s ArticleSelector) SelectArticleByTitleNormalized(tx *ent.Tx, titleNormalized string) (*ent.Article, error) {
	articleEntity, err := tx.Article.Query().
		Where(article.TitleNormalized(titleNormalized)).
		First(s.ctx)
		
	if err != nil {
		return nil, errors.Error(ErrArticleIdNotFound)
	}

	return articleEntity, nil
}