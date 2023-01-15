package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/resetpasswordtoken"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrTokenNotFound ="ErrTokenNotFound: Token not found for this user"
	ErrTokenExpired ="ErrTokenExpired: Token is only valid 2 hours from creation"
)

func (s UserService) GetResetPasswordTokenOwner(token uuid.UUID) (*ent.User, error) {
	rpt, err := database.Client.ResetPasswordToken.Query().Where(resetpasswordtoken.ID(token)).WithOwner().First(s.ctx)
	if err != nil {
		return nil,errors.Error(ErrTokenNotFound)
	}

	twoHoursAfterCreation := rpt.CreatedAt.Add(time.Hour * 2)

	if (time.Now().After(twoHoursAfterCreation)) {
		return nil,errors.Error(ErrTokenExpired)
	}

	return rpt.Edges.Owner, nil
}