package article

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/articletitleguid"
	"github.com/kzmijak/zswod_api_go/ent/image"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrArticleNotFound = "ErrArticleNotFound: Article with given name not found"
)

func (s ArticleService) getArticleEntityByTitle(titleNormalized string, tx *ent.Tx) (*ent.Article, error) {
	articleTitleEntity, err := tx.ArticleTitleGuid.Query().
		Where(articletitleguid.TitleNormalized(titleNormalized)).
		WithArticle(func(aq *ent.ArticleQuery) {
			aq.WithImages(func(iq *ent.ImageQuery) {
				iq.Order(ent.Desc(image.FieldIsPreview)).WithBlob().All(s.ctx)
			}).Only(s.ctx)
		}).
		Only(s.ctx)
	
	if err != nil {
		return nil, errors.Error(ErrArticleNotFound)
	}

	return articleTitleEntity.Edges.Article, nil
}