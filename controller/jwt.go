package controller

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	jwts "github.com/kzmijak/zswod_api_go/services/jwt"
)

func (Controller) ExtractToken(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func (c Controller) ExtractTokenID(ctx *gin.Context) (*uuid.UUID, error) {
	tokenString := c.ExtractToken(ctx)
	token, err := c.jwtService.ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uidString := claims[jwts.CLAIM_GUID].(string)
		
		uuid, err := uuid.Parse(uidString)
		
		if err != nil {
			return nil, err
		}

		return &uuid, nil
	}

	return nil, nil
}