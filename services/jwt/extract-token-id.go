package jwt

import (
	"github.com/google/uuid"
)

func (s JwtService) ExtractTokenID(tokenString string) (*uuid.UUID, error) {
	uuidString, err := s.extractClaim(CLAIM_GUID, tokenString)
	
	if err != nil {
		return nil, err
	}

	uuid, err := uuid.Parse(uuidString)

	if err != nil {
		return nil, err
	}

	return &uuid, nil
}