package jwtManager

import (
	"auth-server/model"
	"auth-server/tokenManager"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type jwtClaims struct {
	jwt.Claims
	tokenManager.TokenData
}

type JwtManager struct {
	secretKey     string
	publicKey     string
	duration      time.Duration
	signingMethod jwt.SigningMethod
}

func (j *JwtManager) createClaims(user model.User) jwtClaims {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(j.duration).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "github.com/simp7/auth-server",
	}
	return jwtClaims{
		Claims: claims,
		TokenData: tokenManager.TokenData{
			Uid:  user.Uid,
			Role: user.Role,
		},
	}
}

func (j *JwtManager) Generate(user model.User) (string, error) {
	token := jwt.NewWithClaims(j.signingMethod, j.createClaims(user))
	block, _ := pem.Decode([]byte(j.secretKey))

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("Error parsing private key", err.Error())
		return "", err
	}
	return token.SignedString(key)
}

func (j *JwtManager) Verify(accessToken string) (tokenManager.TokenData, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			return nil, fmt.Errorf("unexpected token signing method")
		}
		return []byte(j.publicKey), nil
	}
	token, err := jwt.ParseWithClaims(accessToken, &jwtClaims{}, keyFunc)
	if err != nil {
		return tokenManager.TokenData{}, fmt.Errorf("invalid claims: %v", err)
	}
	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return tokenManager.TokenData{}, fmt.Errorf("invalid token")
	}
	return claims.TokenData, nil
}

func NewWithECDSA() (*JwtManager, error) {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate key")
	}
	secretKey, publicKey, err := pemKeyPair(key)
	return &JwtManager{secretKey: string(secretKey), publicKey: string(publicKey), duration: time.Minute * 30, signingMethod: jwt.SigningMethodES256}, nil
}

func pemKeyPair(key *ecdsa.PrivateKey) (privateKey []byte, publicKey []byte, err error) {
	der, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return nil, nil, err
	}

	privateKey = pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: der,
	})

	der, err = x509.MarshalPKIXPublicKey(key.Public())
	if err != nil {
		return nil, nil, err
	}

	publicKey = pem.EncodeToMemory(&pem.Block{
		Type:  "EC PUBLIC KEY",
		Bytes: der,
	})

	return
}
