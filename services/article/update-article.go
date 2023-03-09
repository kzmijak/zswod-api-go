package article

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	articleRepo "github.com/kzmijak/zswod_api_go/repositories/article"
	"github.com/kzmijak/zswod_api_go/services/image"
)


const (
	ErrCouldNotUpdate = "ErrCouldNotUpdate: Failed to update the article"
	ErrCouldNotSanitizeTitle = "ErrCouldNotSanitizeTitle: Failed to sanitize the article title"
)

type ArticleUpdatePayload struct {
	BaseArticlePayload
}

type ImageUpdatePayload struct {
	BaseImagePayload
}

type UpdateArticleRequest struct {
	TitleNormalized string `json:"titleNormalized"`
	Article BaseArticlePayload `json:"article"`
	Images []BaseImagePayload `json:"images" binding:"min=1"`
}

func (s ArticleService) UpdateArticle(req UpdateArticleRequest, tx *ent.Tx) (a *ent.Article, e error) {
	affectedArticle, err := articleRepo.ArticleTitleTx(tx).
		QueryArticleEntityByTitle(req.TitleNormalized).
		Only(s.ctx)

	if err != nil {
		e = err
		return
	}

	if err = affectedArticle.Edges.TitleNormalized.Update().
		SetTitleNormalized(SanitizeTitle(req.Article.Title)).
		Exec(s.ctx);
		err != nil {
		e = errors.Error(ErrCouldNotSanitizeTitle)
		return 
	}

	if err = affectedArticle.Update().SetTitle(req.Article.Title).
		SetShort(req.Article.Short).
		SetContent(req.Article.Content).
		Exec(s.ctx); 
		err != nil {
		e = errors.Error(ErrCouldNotUpdate)
		return
	}

	for _, img := range req.Images {
		s.imageService.CreateImage(image.CreateImageRequest{
			Title: img.Title,
			Alt: img.Alt,
			BlobId: img.BlobId,
			ArticleId: affectedArticle.ID,
			IsPreview: img.IsPreview,
		}, tx)
	}

	a = affectedArticle
	return
}