package storage

import (
	"github.com/simp7/auth-server/model"

	_ "github.com/lib/pq"
)

// User is storage for user.
type User interface {
	FindUser(email string) (model.User, bool)
	GetUser(id model.UserIdentifier) (model.User, error)
	SetUser(user model.User) error
	RemoveUser(id model.UserIdentifier) error
	Close() error
}

// Token is storage for refresh token tied with accessToken and uid.
type Token interface {
	RegisterTokenPair(accessToken string, refreshToken string, uid uint64) error
	DisableToken(refreshToken string) error
	DisableTokenByUid(uid uint64) error
	GetRefreshToken(accessToken string) (string, error)
	Close() error
}

type Storage interface {
	User
	Token
}
