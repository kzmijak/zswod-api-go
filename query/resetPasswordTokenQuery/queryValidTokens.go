package resetPasswordTokenQuery

import (
	"time"

	"github.com/kzmijak/zswod_api_go/ent/resetpasswordtoken"
)

func (rptq ResetPasswordTokenQuery) QueryValidTokens() ResetPasswordTokenQuery {
	query := rptq.Where(resetpasswordtoken.ExpiryTimeGT(time.Now()))
	return FromQuery(query)
}
