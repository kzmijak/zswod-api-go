package resetPasswordTokenQuery

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/resetpasswordtoken"
)

func (rptq ResetPasswordTokenQuery) QueryById(tokenId uuid.UUID) ResetPasswordTokenQuery {
	query := rptq.Where(resetpasswordtoken.ID(tokenId))
	return FromQuery(query)
}