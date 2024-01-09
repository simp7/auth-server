package tokenManager

import (
	"auth-server/model"
)

type TokenData struct {
	Uid  uint64   `json:"uid"`
	Role []string `json:"role"`
}

type TokenManager interface {
	Generate(model.User) (string, error)
	Verify(accessToken string) (TokenData, error)
}
