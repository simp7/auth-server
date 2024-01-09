package jwtManager

import (
	"auth-server/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtClaim struct {
	jwt.Claims
	Uid  uint64   `json:"uid"`
	Role []string `json:"role"`
}

type JwtManager struct {
	secretKey string
	duration  time.Duration
}

func (j *JwtManager) createClaims(user model.User) jwtClaim {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(j.duration).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "github.com/simp7/auth-server",
	}
	return jwtClaim{
		Claims: claims,
		Uid:    user.Uid,
		Role:   user.Role,
	}
}

func (j *JwtManager) Generate(user model.User) (string, error) {
	token := jwt.NewWithClaims(&jwt.SigningMethodECDSA{
		Name:      "",
		Hash:      0,
		KeySize:   0,
		CurveBits: 0,
	}, j.createClaims(user))
	return token.SignedString([]byte(j.secretKey))
}

func (j *JwtManager) Verify(accessToken string) {
	//TODO implement me
	panic("implement me")
}
