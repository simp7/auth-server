package service

import (
	"auth-server/model"
	"auth-server/storage"
	"auth-server/tokenManager"
	"context"
	"github.com/simp7/idl/gen/go/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	auth.UnimplementedAuthServer
	userStorage  storage.User
	tokenStorage storage.Token
	tokenManager tokenManager.TokenManager
}

func NewServer(userStorage storage.User, tokenStorage storage.Token, tokenManager tokenManager.TokenManager) auth.AuthServer {
	s := new(server)
	s.userStorage = userStorage
	s.tokenStorage = tokenStorage
	s.tokenManager = tokenManager
	return s
}

func (s *server) RegisterUser(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	u := model.User{
		Email:    request.Email,
		Password: request.Password,
		Nickname: request.Nickname,
		Role:     []string{"user"},
	}

	if err := s.userStorage.SetUser(u); err != nil {
		return nil, err
	}

	token, err := s.tokenManager.Generate(u)
	if err != nil {
		return nil, err
	}

	if err = s.tokenStorage.RegisterToken(token); err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{Token: token}, nil
}

func (s *server) UnregisterUser(ctx context.Context, request *auth.UnregisterRequest) (*auth.UnregisterResponse, error) {
	data, err := s.tokenManager.Verify(request.Token)
	if err != nil {
		return nil, err
	}
	if err = s.userStorage.RemoveUser(model.UserIdentifier{Uid: data.Uid}); err != nil {
		return nil, err
	}
	return &auth.UnregisterResponse{Success: true}, nil
}

func (s *server) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	switch data := request.Method.(type) {
	case *auth.LoginRequest_Traditional:
		u, ok := s.userStorage.FindUser(data.Traditional.Email)
		if !ok {
			return nil, status.Errorf(codes.NotFound, "user not found: %v", data.Traditional.Email)
		}

		if err := u.VerifyPassword(data.Traditional.Password); err != nil {
			return nil, status.Error(codes.PermissionDenied, "incorrect password")
		}

		token, err := s.tokenManager.Generate(u)
		if err != nil {
			return nil, err
		}

		err = s.tokenStorage.RegisterToken(token)
		if err != nil {
			return nil, err
		}

		return &auth.LoginResponse{Token: token}, nil
	case *auth.LoginRequest_Oauth:
		return nil, status.Error(codes.Unimplemented, "oauth is not implemented yet")
	}
	return nil, status.Error(codes.Internal, "could not recognize method for login")
}

func (s *server) Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	err := s.tokenStorage.UnregisterToken(request.Token)
	if err != nil {
		return nil, err
	}
	return &auth.LogoutResponse{Token: request.Token}, nil
}

func (s *server) mustEmbedUnimplementedAuthServer() {
}
