package storage

import "auth-server/model"

type Storage interface {
	FindUser(email string) (model.User, bool)
	GetUser(id model.UserIdentifier) (model.User, error)
	SetUser(user model.User) error
	RemoveUser(id model.UserIdentifier) error
}
