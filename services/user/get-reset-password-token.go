package user

import (
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrCouldNotParseResetPasswordTokenId = "ErrCouldNotParseResetPasswordTokenId: Could not parse specified string to UUID"
	ErrResetPasswordTokenIdNotFound = "ErrResetPasswordTokenIdNotFound: Could not find resetPasswordToken by given ID"
)

func (s UserService) GetResetPasswordToken(tokenId string, tx *ent.Tx) (*ent.ResetPasswordToken, error) {
	tokenIdParsed, err := uuid.Parse(tokenId)
	if err != nil {
		return nil, errors.Error(ErrCouldNotParseUserId)
	}

	user, err := tx.ResetPasswordToken.Get(s.ctx, tokenIdParsed)
	if err != nil {
		return nil, errors.Error(ErrUserIdNotFound)
	}

	return user, nil
} 