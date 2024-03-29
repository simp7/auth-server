package jwtManager

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/simp7/auth-server/model"
	"github.com/simp7/auth-server/tokenManager"
	"time"
)

const (
	AccessTokenDuration  = time.Hour
	RefreshTokenDuration = 60 * 24 * time.Hour
)

type jwtClaims struct {
	jwt.Claims
	tokenManager.TokenData
}

type manager[T jwt.SigningMethod] struct {
	secretKey     string
	publicKey     string
	signingMethod T
}

func (m *manager[T]) createClaims(user model.User, duration time.Duration) jwtClaims {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
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

func (m *manager[T]) tokenize(user model.User, duration time.Duration) (string, error) {
	refreshToken := jwt.NewWithClaims(m.signingMethod, m.createClaims(user, duration))
	block, _ := pem.Decode([]byte(m.secretKey))
	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("error parsing private key: %v", err.Error())
	}
	return refreshToken.SignedString(key)
}

func (m *manager[T]) Generate(user model.User) (tokenManager.Tokens, error) {
	accessToken, err := m.tokenize(user, AccessTokenDuration)
	if err != nil {
		return tokenManager.Tokens{}, fmt.Errorf("error parsing private key: %v", err.Error())
	}

	refreshToken, err := m.tokenize(user, RefreshTokenDuration)
	if err != nil {
		return tokenManager.Tokens{}, fmt.Errorf("error parsing private key: %v", err.Error())
	}

	return tokenManager.Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func (m *manager[T]) Verify(accessToken string) (tokenManager.TokenData, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(T)
		if !ok {
			return nil, fmt.Errorf("unexpected token signing method")
		}
		return []byte(m.publicKey), nil
	}
	token, err := jwt.ParseWithClaims(accessToken, &jwtClaims{}, keyFunc, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return tokenManager.TokenData{}, fmt.Errorf("invalid claims: %v", err)
	}
	claims, ok := token.Claims.(*jwtClaims)
	if !ok {
		return tokenManager.TokenData{}, fmt.Errorf("invalid token")
	}

	return claims.TokenData, nil
}

func ECDSA(key *ecdsa.PrivateKey) (tokenManager.TokenManager, error) {
	secretKey, publicKey, err := ecdsaKeyPair(key)
	if err != nil {
		return nil, fmt.Errorf("failed to generate key pair: %v", err)
	}

	return &manager[*jwt.SigningMethodECDSA]{secretKey: string(secretKey), publicKey: string(publicKey), signingMethod: jwt.SigningMethodES256}, nil
}

func ecdsaKeyPair(key *ecdsa.PrivateKey) (privateKey []byte, publicKey []byte, err error) {
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
