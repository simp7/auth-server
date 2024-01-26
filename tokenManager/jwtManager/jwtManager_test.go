package jwtManager

import (
	"auth-server/model"
	"auth-server/tokenManager"
	"crypto/ecdsa"
	"github.com/golang-jwt/jwt/v5"
	"reflect"
	"testing"
	"time"
)

func TestECDSA(t *testing.T) {
	type args struct {
		key *ecdsa.PrivateKey
	}
	tests := []struct {
		name    string
		args    args
		want    tokenManager.TokenManager
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ECDSA(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ECDSA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ECDSA() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ecdsaKeyPair(t *testing.T) {
	type args struct {
		key *ecdsa.PrivateKey
	}
	tests := []struct {
		name           string
		args           args
		wantPrivateKey []byte
		wantPublicKey  []byte
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrivateKey, gotPublicKey, err := ecdsaKeyPair(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ecdsaKeyPair() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPrivateKey, tt.wantPrivateKey) {
				t.Errorf("ecdsaKeyPair() gotPrivateKey = %v, want %v", gotPrivateKey, tt.wantPrivateKey)
			}
			if !reflect.DeepEqual(gotPublicKey, tt.wantPublicKey) {
				t.Errorf("ecdsaKeyPair() gotPublicKey = %v, want %v", gotPublicKey, tt.wantPublicKey)
			}
		})
	}
}

func Test_manager_Generate(t *testing.T) {
	type args struct {
		user model.User
	}
	type testCase[T jwt.SigningMethod] struct {
		name    string
		m       manager[T]
		args    args
		want    tokenManager.Tokens
		wantErr bool
	}
	tests := []testCase[*jwt.SigningMethodECDSA]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Generate(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Generate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_manager_Verify(t *testing.T) {
	type args struct {
		accessToken string
	}
	type testCase[T jwt.SigningMethod] struct {
		name    string
		m       manager[T]
		args    args
		want    tokenManager.TokenData
		wantErr bool
	}
	tests := []testCase[*jwt.SigningMethodECDSA]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Verify(tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Verify() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_manager_createClaims(t *testing.T) {
	type args struct {
		user model.User
	}
	type testCase[T jwt.SigningMethod] struct {
		name string
		m    manager[T]
		args args
		want jwtClaims
	}
	tests := []testCase[*jwt.SigningMethodECDSA]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.createClaims(tt.args.user, time.Second); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createClaims() = %v, want %v", got, tt.want)
			}
		})
	}
}
