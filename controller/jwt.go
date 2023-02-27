package controller

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
		if value, ok := claims[claimKey].(string); ok {
			return value, nil
		} 
		if value, ok := claims[claimKey].(float64); ok {
			str := fmt.Sprintf("%.0f", value)
    	return strings.TrimRight(str, ".0"), nil
		}
		
		return "", nil
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

func (c Controller) CheckValid(ctx *gin.Context) bool {
	expString, err := c.extractClaim(jwts.CLAIM_EXPIRED, ctx)

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