package customPageService

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/customPageModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/utils/sanitizer"
)

const (
	ErrCouldNotCreateCustomPage = "ErrCouldNotCreateCustomPage: Failed to create custom page"
)

func (s CustomPageService) CreateCustomPage(req customPageModel.CreateCustomPagePayload, tx *ent.Tx) (customPageId uuid.UUID, err error) {
	titleSanitized := sanitizer.SanitizeString(req.Title)	
	sectionSanitized := sanitizer.SanitizeString(req.Section)	
	url := sectionSanitized + "/" + titleSanitized 
	order, err := s.selectors.SelectSectionMembersCount(tx, req.Section)

	if err != nil {
		return uuid.Nil, err
	}
	
	customPageId = uuid.New()
	err = tx.CustomPage.Create().
		SetID(customPageId).
		SetOrder(order).
		SetTitleNormalized(url).
		SetSection(req.Section).
		SetIconId(req.IconId).
		SetTitle(req.Title).
		SetContent(req.Content).
		Exec(s.ctx)
	
	if err != nil {
		return uuid.Nil, errors.Error(ErrCouldNotCreateCustomPage)
	}

	return customPageId, nil
}
