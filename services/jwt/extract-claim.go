package jwt

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)


func (c JwtService) extractClaim(claimKey string, tokenString string) (string, error) {
	token, err := c.ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		if value, ok := claims[claimKey].(string); ok {
			return value, nil
		} 
		if value, ok := claims[claimKey].(float64); ok {
			if value == 0 {
				return "0", nil;
			} 

			str := fmt.Sprintf("%.0f", value)
    	return strings.TrimRight(str, ".0"), nil
		}
		
		return "", nil
	}

	return "", nil
}