package inMemory

import (
	"auth-server/model"
	"auth-server/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type inMemory struct {
	userById        map[model.UserIdentifier]model.User
	userByEmail     map[string]model.User
	accessToRefresh map[string]string
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
	u, err := i.GetUser(id)
	if err != nil {
		return err
	}
	delete(i.userById, id)
	delete(i.userByEmail, u.Email)
	return nil
}

func (i *inMemory) RegisterTokenPair(accessToken string, refreshToken string) error {
	i.accessToRefresh[accessToken] = refreshToken
	return nil
}

func (i *inMemory) GetRefreshToken(accessToken string) (string, error) {
	refreshToken, ok := i.accessToRefresh[accessToken]
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "token is not valid: %v", accessToken)
	}
	return refreshToken, nil
}

func (i *inMemory) DisableToken(token string) error {
	if _, ok := i.accessToRefresh[token]; !ok {
		return status.Errorf(codes.NotFound, "token not exist: %v", token)
	}
	delete(i.accessToRefresh, token)
	return nil
}

func (i *inMemory) Close() error {
	i.accessToRefresh = nil
	i.userById = nil
	i.userByEmail = nil
	return nil
}

func Storage() storage.Storage {
	s := new(inMemory)
	s.accessToRefresh = make(map[string]string)
	s.userById = make(map[model.UserIdentifier]model.User)
	s.userByEmail = make(map[string]model.User)

	return s
}
