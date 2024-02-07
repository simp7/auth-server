package service

import (
	"context"
	"github.com/simp7/auth-server/model"
	"github.com/simp7/auth-server/storage"
	"github.com/simp7/auth-server/tokenManager"
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
	index        uint64
}

// NewServer returns new AuthServer struct.
// user storage and token storage can be different: recommended composition is Redis and RDB(SQL Family).
func NewServer(userStorage storage.User, tokenStorage storage.Token, tokenManager tokenManager.TokenManager) auth.AuthServer {
	s := new(server)
	s.userStorage = userStorage
	s.tokenStorage = tokenStorage
	s.tokenManager = tokenManager
	s.index = 1
	return s
}

func (s *server) getIndex() uint64 {
	result := s.index
	s.index = s.index + 1
	return result
}

// RegisterUser is request for creating new user.
// This also includes generating-password process.
func (s *server) RegisterUser(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "error when generating hashed password")
	}
	u := model.User{
		UserIdentifier: model.UserIdentifier{
			Uid: s.getIndex(),
		},
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

	if err = s.tokenStorage.RegisterTokenPair(token.Access, token.Refresh, u.Uid); err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{AccessToken: token.Access}, nil
}

// UnregisterUser is request for deleting user from db.
// This function compares token and requested uid, for preventing deleting other.
func (s *server) UnregisterUser(ctx context.Context, request *auth.UnregisterRequest) (*auth.UnregisterResponse, error) {
	token, err := s.getTokenFromMetadata(ctx)
	if err != nil {
		return nil, err
	}

	extracted, err := s.tokenManager.Verify(token)
	if err != nil {
		return nil, err
	}
	if extracted.Uid != request.Uid {
		return nil, status.Error(codes.PermissionDenied, "cannot permit unregistering by other")
	}

	if err = s.userStorage.RemoveUser(model.UserIdentifier{Uid: request.Uid}); err != nil {
		return nil, err
	}
	return &auth.UnregisterResponse{Success: true}, nil
}

func (s *server) refreshTokenRotation() (accessToken string, refreshToken string) {
	return
}

// Login is request for getting access token and refresh token of user.
// This function also invalidates previous refresh token and create new token: refresh-token-rotation
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

		if err := s.tokenStorage.DisableTokenByUid(u.Uid); err != nil {
			return nil, err
		}

		token, err := s.tokenManager.Generate(u)
		if err != nil {
			return nil, err
		}

		if err = s.tokenStorage.RegisterTokenPair(token.Access, token.Refresh, u.Uid); err != nil {
			return nil, err
		}

		return &auth.LoginResponse{AccessToken: token.Access}, nil
	case *auth.LoginRequest_Oauth:
		return nil, status.Error(codes.Unimplemented, "oauth is not implemented yet")
	}
	return nil, status.Error(codes.Internal, "could not recognize method for login")
}

// Logout is request for discarding previous refresh token of user.
func (s *server) Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	token := request.RefreshToken
	if err := s.tokenStorage.DisableToken(token); err != nil {
		return nil, err
	}
	return &auth.LogoutResponse{RefreshToken: request.RefreshToken}, nil
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
