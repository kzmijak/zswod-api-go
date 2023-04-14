package customPageService

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/customPageModel"
)

func (s CustomPageService) GetCustomPageByTitleNormalized(url string, tx *ent.Tx) (customPageModel.CustomPageModel, error) {
	return s.selectors.SelectCustomPageByUrl(tx, url)
}