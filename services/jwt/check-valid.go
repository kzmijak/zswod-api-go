package jwt

import (
	"strconv"
	"time"
)

func (s JwtService) CheckValid(token string) bool {
	expString, err := s.extractClaim(CLAIM_EXPIRED, token)

	if err != nil {
		return false
	}

	exp, err := strconv.Atoi(expString)
	if err != nil {
		return false
	}

	now := int(time.Now().Unix())

	return exp > now
}