package user

import (
	"github.com/kzmijak/zswod_api_go/ent/user"
	"github.com/kzmijak/zswod_api_go/models/role"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	"golang.org/x/crypto/bcrypt"
)

type SignInRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

const (
	ErrUserNotFound = "ErrUserNotFound: Could not match any user with provided username"
	ErrInvalidPassword = "ErrInvalidPassword: Password for the user is incorrect"
	ErrInvalidClaims = "ErrInvalidClaims: Could not generate token for this user"
	ErrNoRoles = "ErrNoRoles: User has no roles and could not be retrieved"
	ErrUnknownRole = "ErrUnknownRole: User has unknown roles and could not be retrieved"
)

func (s *UserService) SignIn(request SignInRequest) (string, error) {
	user, err := database.Client.User.Query().Where(user.Email(request.Email)).Only(s.ctx);

	if err != nil {
		return "" ,errors.Error(ErrUserNotFound)
	}
	
	userRole, err := user.QueryRoles().Only(s.ctx)
	if err != nil {
		return "", errors.Error(ErrNoRoles)
	}

	role, exists := role.FromString(userRole.ID)
	if !exists {
		return "", errors.Error(ErrUnknownRole)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return "", errors.Error(ErrInvalidPassword)
	}

	jwt, err := s.jwtService.GenerateToken(string(user.ID.String()), role)

	if err != nil {
		return "", errors.Error(ErrInvalidClaims)
	}

	return jwt, nil;
}