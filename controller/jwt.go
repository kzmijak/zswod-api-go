package controller

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	jwts "github.com/kzmijak/zswod_api_go/services/jwt"
)

func (Controller) ExtractToken(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func (c Controller) ExtractTokenID(ctx *gin.Context) (uint, error) {
	tokenString := c.ExtractToken(ctx)
	token, err := c.jwtService.ParseToken(tokenString)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims[jwts.CLAIM_GUID]), 10, 32)
		if err != nil {
			return 0, err
		}

		return uint(uid), nil
	}

	return 0, nil
}