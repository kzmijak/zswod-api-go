package resetPasswordTokenQuery

import "github.com/kzmijak/zswod_api_go/ent"

type ResetPasswordTokenQuery struct {
	*ent.ResetPasswordTokenQuery
}

func FromQuery(query *ent.ResetPasswordTokenQuery) ResetPasswordTokenQuery {
	return ResetPasswordTokenQuery{
		query,
	}
}

func FromTx (tx *ent.Tx) ResetPasswordTokenQuery {
	return FromQuery(tx.ResetPasswordToken.Query())
}