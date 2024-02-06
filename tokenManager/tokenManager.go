package tokenManager

import (
	"github.com/simp7/auth-server/model"
)

type TokenData struct {
	Uid  uint64   `json:"uid"`
	Role []string `json:"role"`
}

type Tokens struct {
	Access  string
	Refresh string
}

type TokenManager interface {
	Generate(model.User) (Tokens, error)
	Verify(accessToken string) (TokenData, error)
}
