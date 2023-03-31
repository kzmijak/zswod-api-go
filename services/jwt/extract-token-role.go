package jwt

import (
	"strconv"

	"github.com/kzmijak/zswod_api_go/models/role"
	"github.com/kzmijak/zswod_api_go/modules/errors"
)

const (
	ErrNoRole = "ErrNoRole: User role is invalid"
)

func (c JwtService) ExtractTokenRole(tokenString string) (*role.Role, error) {
	roleString, err := c.extractClaim(CLAIM_ROLE, tokenString)
	
	if err != nil {
		return nil, err
	}

	roleId, err := strconv.Atoi(roleString)
	if err != nil {
		return nil, errors.Error(ErrNoRole)
	}

	role, err := role.FromId(roleId)
	if err != nil {
		return nil, err
	}
	
	return &role, nil
}