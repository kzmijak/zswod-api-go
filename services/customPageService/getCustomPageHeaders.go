package customPageService

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/customPageHeaderModel"
)

func (s CustomPageService) GetCustomPageHeaders(tx *ent.Tx) ([]customPageHeaderModel.CustomPageHeaderModel, error) {
	customPageEntities, err := s.selectors.SelectAllCustomPages(tx)
	if err != nil {
		return nil, err
	}

	return customPageHeaderModel.ArrayFromEntities(customPageEntities)
}