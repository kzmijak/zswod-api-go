package user

import (
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/user"
)

const (
	ErrUserEmailNotFound = "ErrUserEmailNotFound: Could not find user with given email"
)

func (s UserService) GetUserByEmail(email string, tx *ent.Tx) (*ent.User, error) {
	user, err := tx.User.Query().Where(user.Email(email)).Only(s.ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}