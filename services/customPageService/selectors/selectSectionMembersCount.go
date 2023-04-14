package selectors

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"github.com/kzmijak/zswod_api_go/query/customPageQuery"
)

const (
	ErrCouldNotCountCustomPagesBySection = "ErrCouldNotCountCustomPagesBySection: Failed to count custom pages by section"
)

func (s CustomPageSelector) SelectSectionMembersCount(tx *ent.Tx, section string) (count int, err error) {
	count, err = customPageQuery.FromTx(tx).
		QueryBySection(section).
		Count(s.ctx)

	if err != nil {
		return 0, errors.Error(ErrCouldNotCountCustomPagesBySection)
	}

	return count, nil
}