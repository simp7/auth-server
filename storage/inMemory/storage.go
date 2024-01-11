package inMemory

import (
	"auth-server/model"
	"auth-server/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type inMemory struct {
	userById    map[model.UserIdentifier]model.User
	userByEmail map[string]model.User
	validToken  map[string]struct{}
}

func (i *inMemory) FindUser(email string) (model.User, bool) {
	user, ok := i.userByEmail[email]
	return user, ok
}

func (i *inMemory) GetUser(id model.UserIdentifier) (model.User, error) {
	user, ok := i.userById[id]
	if !ok {
		return user, status.Errorf(codes.NotFound, "user not found: %v", id)
	}
	return user, nil
}

func (i *inMemory) SetUser(user model.User) error {
	_, ok := i.userById[user.UserIdentifier]
	if ok {
		return status.Errorf(codes.AlreadyExists, "user already exist: %v", user.Email)
	}
	i.userById[user.UserIdentifier] = user
	i.userByEmail[user.Email] = user
	return nil
}

func (i *inMemory) RemoveUser(id model.UserIdentifier) error {
	if _, err := i.GetUser(id); err != nil {
		return err
	}
	delete(i.userById, id)
	return nil
}

func (i *inMemory) RegisterToken(token string) error {
	if _, ok := i.validToken[token]; ok {
		return status.Errorf(codes.AlreadyExists, "token already exist: %v", token)
	}

	i.validToken[token] = struct{}{}
	return nil
}

func (i *inMemory) CheckToken(token string) error {
	_, ok := i.validToken[token]
	if !ok {
		return status.Errorf(codes.Unauthenticated, "token is not valid: %v", token)
	}
	return nil
}

func (i *inMemory) UnregisterToken(token string) error {
	if _, ok := i.validToken[token]; !ok {
		return status.Errorf(codes.NotFound, "token not exist: %v", token)
	}
	delete(i.validToken, token)
	return nil
}

func Storage() storage.Storage {
	s := new(inMemory)
	s.validToken = make(map[string]struct{})
	s.userById = make(map[model.UserIdentifier]model.User)
	s.userByEmail = make(map[string]model.User)

	return s
}
