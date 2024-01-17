package service

import (
	"auth-server/model"
	"auth-server/storage"
	"auth-server/tokenManager"
	"context"
	"github.com/simp7/idl/gen/go/auth"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "error when generate hashed password")
	}
	u := model.User{
		Email:    request.Email,
		Password: string(hashedPassword),
		Nickname: request.Nickname,
		Role:     []string{"user"},
	}

	if err = s.userStorage.SetUser(u); err != nil {
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
	if err := s.userStorage.RemoveUser(model.UserIdentifier{Uid: request.Uid}); err != nil {
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
	token := request.Token
	meta, err := s.getTokenFromMetadata(ctx)
	if err != nil {
		return nil, err
	}
	if token != meta {
		return nil, status.Error(codes.PermissionDenied, "argument and current user are not matching")
	}
	err = s.tokenStorage.UnregisterToken(request.Token)
	if err != nil {
		return nil, err
	}
	return &auth.LogoutResponse{Token: request.Token}, nil
}

func (s *server) mustEmbedUnimplementedAuthServer() {
}

func (s *server) getTokenFromMetadata(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "token is not provided")
	}
	t, ok := md["authorization"]
	if !ok {
		return "", status.Error(codes.Unauthenticated, "token is not provided")
	}
	return t[0], nil
}
