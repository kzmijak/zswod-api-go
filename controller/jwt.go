package controller

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/models/role"
	"github.com/kzmijak/zswod_api_go/modules/errors"
	jwts "github.com/kzmijak/zswod_api_go/services/jwt"
)

const (
	ErrNoRole = "ErrNoRole: User role is invalid"
)

func (Controller) ExtractToken(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func (c Controller) extractClaim(claimKey string, ctx *gin.Context) (string, error) {
	tokenString := c.ExtractToken(ctx)
	token, err := c.jwtService.ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		claim := claims[claimKey].(string) // TODO: FIX FOR ROLES!!!
		
		return claim, nil
	}

	return "", nil
}

func (c Controller) ExtractTokenID(ctx *gin.Context) (*uuid.UUID, error) {
	uuidString, err := c.extractClaim(jwts.CLAIM_GUID, ctx)
	
	if err != nil {
		return nil, err
	}

	uuid, err := uuid.Parse(uuidString)

	if err != nil {
		return nil, err
	}

	return &uuid, nil
}

func (c Controller) ExtractTokenRole(ctx *gin.Context) (*role.Role, error) {
	roleString, err := c.extractClaim(jwts.CLAIM_ROLE, ctx)
	
	if err != nil {
		return nil, err
	}

	role, exists := role.FromString(roleString)

	if !exists {
		return nil, errors.Error(ErrNoRole)
	}
	
	return &role, nil
}