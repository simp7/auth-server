package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_VerifyPassword(t *testing.T) {

	type fields struct {
		UserIdentifier UserIdentifier
		Email          string
		Password       string
		Nickname       string
		Role           []string
		uid            int
	}
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test1",
			fields: fields{
				UserIdentifier: UserIdentifier{1},
				Email:          "tester1@test.com",
				Password:       "pass1234",
				Nickname:       "test",
				Role:           nil,
			},
			args: args{
				password: "password",
			},
			wantErr: true,
		},
		{
			name: "Test2",
			fields: fields{
				UserIdentifier: UserIdentifier{1},
				Email:          "tester1@test.com",
				Password:       "pass1234",
				Nickname:       "test",
				Role:           nil,
			},
			args: args{
				password: "pass1234",
			},
			wantErr: true,
		},
		{
			name: "Test3",
			fields: fields{
				UserIdentifier: UserIdentifier{1},
				Email:          "tester1@test.com",
				//Password:       string(hashedPassword),
				Password: "$2a$10$F55KexLrQOs5mY9X6hQJ5OgpS3J1P27l9.ZJLVaxA8wzAKM305wC6",
				Nickname: "test",
				Role:     nil,
			},
			args: args{
				password: "pass1234",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				UserIdentifier: tt.fields.UserIdentifier,
				Email:          tt.fields.Email,
				Password:       tt.fields.Password,
				Nickname:       tt.fields.Nickname,
				Role:           tt.fields.Role,
			}
			err := u.VerifyPassword(tt.args.password)
			if tt.wantErr {
				assert.Errorf(t, err, "want error, got no error\n")
			} else {
				assert.NoErrorf(t, err, "do not want any error, got error: %v\n", err)
			}
		})
	}
}
