package resetPasswordTokenQuery

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/query/userQuery"
)

func (rptq ResetPasswordTokenQuery) QueryFullResetPasswordToken() ResetPasswordTokenQuery {
	query := rptq.QueryValidTokens().WithOwner(
		func(uq *ent.UserQuery) { userQuery.FromQuery(uq).QueryFullUser() },
	)
	return FromQuery(query)
}