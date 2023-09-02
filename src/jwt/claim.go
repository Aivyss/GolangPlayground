package jwt

import (
	"com.playground/dto"
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

type ClaimData struct {
	UserId   int  `json:"userId"`
	Verified bool `json:"verified"`
}

type CustomClaim struct {
	ClaimData `json:"data"`
	jwt.StandardClaims
}

func (c ClaimData) GenerateTokens() (access string, refresh string, err error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaim{
		c,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(accessLifeTime).UnixMilli(),
			Issuer:    "test",
		},
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaim{
		c,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshLifeTime).UnixMilli(),
			Issuer:    "test",
		},
	})

	access, err1 := accessToken.SignedString(accessSecretKey)
	refresh, err2 := refreshToken.SignedString(refreshSecretKey)

	if err1 != nil || err2 != nil {
		access = ""
		refresh = ""

		return access, refresh, errors.New("fail to create jwt")
	} else {
		return access, refresh, nil
	}
}

func ParseAccessTokenString(tokenString string) (ClaimData, error) {
	data := new(CustomClaim)
	token, err := jwt.ParseWithClaims(tokenString, data, func(token *jwt.Token) (interface{}, error) {
		return accessSecretKey, nil
	})

	result := ClaimData{}
	if err == nil {
		if claim, ok := token.Claims.(*CustomClaim); ok {
			result = claim.ClaimData

			return result, nil
		} else {
			return result, errors.New("fail to parse access token string")
		}
	} else {
		return result, errors.New("fail to parse access token string")
	}
}

func ParseRefreshTokenString(tokenString string) (ClaimData, error) {
	data := new(CustomClaim)
	token, err := jwt.ParseWithClaims(tokenString, data, func(token *jwt.Token) (interface{}, error) {
		return refreshSecretKey, nil
	})

	result := ClaimData{}
	if err != nil {
		if claim, ok := token.Claims.(CustomClaim); ok {
			result = claim.ClaimData

			return result, nil
		} else {
			return result, errors.New("fail to parse refresh token string")
		}
	} else {
		return result, errors.New("fail to parse refresh token string")
	}
}

func NewClaimData(account dto.Account) ClaimData {
	return ClaimData{
		UserId:   account.Id,
		Verified: true,
	}
}
