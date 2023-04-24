package customPageService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
)

func (s CustomPageService) GetCustomPageUrlById(customPageId uuid.UUID, tx *ent.Tx) (string, error) {
	model, err := s.selectors.SelectCustomPageById(tx, customPageId)
	if err != nil {
		return "", err
	}

	return model.Url, nil
}