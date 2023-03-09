package jwt

import (
	"strconv"

	"github.com/kzmijak/zswod_api_go/models/role"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrNoRole = "ErrNoRole: User role is invalid"
)

// TODO: Remember that tokenString is string, whilst the target is to have it as byte array
func (c JwtService) ExtractTokenRole(tokenString string) (*role.Role, error) {
	roleString, err := c.extractClaim(CLAIM_ROLE, tokenString)
	
	if err != nil {
		return nil, err
	}

	roleId, err := strconv.Atoi(roleString)
	if err != nil {
		return nil, errors.Error(ErrNoRole)
	}

	role, exists := role.FromId(roleId)
	if exists {
		return &role, nil
	}
	
	return nil, errors.Error(ErrNoRole)
}