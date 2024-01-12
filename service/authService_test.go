package service

import (
	"auth-server/storage"
	"auth-server/storage/inMemory"
	"auth-server/tokenManager"
	"auth-server/tokenManager/jwtManager"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"github.com/simp7/idl/gen/go/auth"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewServer(t *testing.T) {
	type args struct {
		userStorage  storage.User
		tokenStorage storage.Token
		tokenManager tokenManager.TokenManager
	}
	tests := []struct {
		name string
		args args
		want auth.AuthServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.userStorage, tt.args.tokenStorage, tt.args.tokenManager); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_Login(t *testing.T) {
	s := inMemory.Storage()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err, "error while generating key")
	manager, err := jwtManager.ECDSA(key)
	assert.NoError(t, err, "error while call manager")
	NewServer(s, s, manager)
	type fields struct {
		UnimplementedAuthServer auth.UnimplementedAuthServer
		userStorage             storage.User
		tokenStorage            storage.Token
		tokenManager            tokenManager.TokenManager
	}
	type args struct {
		ctx     context.Context
		request *auth.LoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *auth.LoginResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedAuthServer: tt.fields.UnimplementedAuthServer,
				userStorage:             tt.fields.userStorage,
				tokenStorage:            tt.fields.tokenStorage,
				tokenManager:            tt.fields.tokenManager,
			}
			got, err := s.Login(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_Logout(t *testing.T) {
	type fields struct {
		UnimplementedAuthServer auth.UnimplementedAuthServer
		userStorage             storage.User
		tokenStorage            storage.Token
		tokenManager            tokenManager.TokenManager
	}
	type args struct {
		ctx     context.Context
		request *auth.LogoutRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *auth.LogoutResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedAuthServer: tt.fields.UnimplementedAuthServer,
				userStorage:             tt.fields.userStorage,
				tokenStorage:            tt.fields.tokenStorage,
				tokenManager:            tt.fields.tokenManager,
			}
			got, err := s.Logout(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Logout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Logout() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_RegisterUser(t *testing.T) {
	type fields struct {
		UnimplementedAuthServer auth.UnimplementedAuthServer
		userStorage             storage.User
		tokenStorage            storage.Token
		tokenManager            tokenManager.TokenManager
	}
	type args struct {
		ctx     context.Context
		request *auth.RegisterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *auth.RegisterResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedAuthServer: tt.fields.UnimplementedAuthServer,
				userStorage:             tt.fields.userStorage,
				tokenStorage:            tt.fields.tokenStorage,
				tokenManager:            tt.fields.tokenManager,
			}
			got, err := s.RegisterUser(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_UnregisterUser(t *testing.T) {
	type fields struct {
		UnimplementedAuthServer auth.UnimplementedAuthServer
		userStorage             storage.User
		tokenStorage            storage.Token
		tokenManager            tokenManager.TokenManager
	}
	type args struct {
		ctx     context.Context
		request *auth.UnregisterRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *auth.UnregisterResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedAuthServer: tt.fields.UnimplementedAuthServer,
				userStorage:             tt.fields.userStorage,
				tokenStorage:            tt.fields.tokenStorage,
				tokenManager:            tt.fields.tokenManager,
			}
			got, err := s.UnregisterUser(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnregisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnregisterUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_server_mustEmbedUnimplementedAuthServer(t *testing.T) {
	type fields struct {
		UnimplementedAuthServer auth.UnimplementedAuthServer
		userStorage             storage.User
		tokenStorage            storage.Token
		tokenManager            tokenManager.TokenManager
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedAuthServer: tt.fields.UnimplementedAuthServer,
				userStorage:             tt.fields.userStorage,
				tokenStorage:            tt.fields.tokenStorage,
				tokenManager:            tt.fields.tokenManager,
			}
			s.mustEmbedUnimplementedAuthServer()
		})
	}
}
