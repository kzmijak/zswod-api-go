package jwt

import (
	"github.com/kzmijak/zswod_api_go/models/role"
)

const (
	ErrNoRole = "ErrNoRole: User role is invalid"
)

func (c JwtService) ExtractTokenRole(tokenString string) (*role.Role, error) {
	roleString, err := c.extractClaim(CLAIM_ROLE, tokenString)
	
	if err != nil {
		return nil, err
	}

	role, err := role.FromString(roleString)
	if err != nil {
		return nil, err
	}
	
	return &role, nil
}