package storage

import (
	"auth-server/model"

	_ "github.com/lib/pq"
)

type User interface {
	FindUser(email string) (model.User, bool)
	GetUser(id model.UserIdentifier) (model.User, error)
	SetUser(user model.User) error
	RemoveUser(id model.UserIdentifier) error
	Close() error
}

type Token interface {
	RegisterToken(token string) error
	UnregisterToken(token string) error
	CheckToken(token string) error
	Close() error
}

type Storage interface {
	User
	Token
}
