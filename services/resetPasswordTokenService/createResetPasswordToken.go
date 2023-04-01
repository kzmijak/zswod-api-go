package resetPasswordTokenService

import (
	"time"

	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/resetPasswordTokenModel"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrTokenAlreadyExists = "ErrTokenAlreadyExists: User already requested token creation"
	ErrTokenCreationFailed = "ErrTokenCreationFailed: Failed to create a token for specified user"
)

func (s ResetPasswordTokenService) CreateResetPasswordToken(ownerEmail string, tx *ent.Tx) (resetPasswordTokenModel.ResetPasswordTokenModel, error) {
	existingToken, _ := s.tokenSelectors.SelectTokenByOwnerEmail(tx, ownerEmail)
	if existingToken != nil {
		s.InvalidateResetPasswordToken(existingToken.ID, tx)
	}

	tokenOwner, err := s.userSelectors.SelectUserByEmail(tx, ownerEmail)
	if err != nil {
		return resetPasswordTokenModel.Nil, err
	}

	now := time.Now()
	twoHoursFromNow := now.Add(time.Hour * 2)

	tokenEntity, err := tx.ResetPasswordToken.Create().
		SetCreateTime(now).
		SetExpiryTime(twoHoursFromNow).
		SetID(uuid.New()).
		SetOwner(tokenOwner).
		Save(s.ctx)
	if err != nil {
		return resetPasswordTokenModel.Nil, errors.Error(ErrTokenCreationFailed)
	}


	return resetPasswordTokenModel.FromEntity(tokenEntity)
}


