package inMemory

import (
	"auth-server/model"
	"auth-server/storage"
	"reflect"
	"testing"
)

func TestStorage(t *testing.T) {
	tests := []struct {
		name string
		want storage.Storage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Storage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemory_CheckToken(t *testing.T) {
	type fields struct {
		userById    map[model.UserIdentifier]model.User
		userByEmail map[string]model.User
		validToken  map[string]struct{}
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inMemory{
				userById:    tt.fields.userById,
				userByEmail: tt.fields.userByEmail,
				validToken:  tt.fields.validToken,
			}
			if err := i.CheckToken(tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("CheckToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_inMemory_FindUser(t *testing.T) {
	type fields struct {
		userById    map[model.UserIdentifier]model.User
		userByEmail map[string]model.User
		validToken  map[string]struct{}
	}
	type args struct {
		email string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   model.User
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inMemory{
				userById:    tt.fields.userById,
				userByEmail: tt.fields.userByEmail,
				validToken:  tt.fields.validToken,
			}
			got, got1 := i.FindUser(tt.args.email)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUser() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_inMemory_GetUser(t *testing.T) {
	type fields struct {
		userById    map[model.UserIdentifier]model.User
		userByEmail map[string]model.User
		validToken  map[string]struct{}
	}
	type args struct {
		id model.UserIdentifier
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inMemory{
				userById:    tt.fields.userById,
				userByEmail: tt.fields.userByEmail,
				validToken:  tt.fields.validToken,
			}
			got, err := i.GetUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inMemory_RegisterToken(t *testing.T) {
	type fields struct {
		userById    map[model.UserIdentifier]model.User
		userByEmail map[string]model.User
		validToken  map[string]struct{}
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inMemory{
				userById:    tt.fields.userById,
				userByEmail: tt.fields.userByEmail,
				validToken:  tt.fields.validToken,
			}
			if err := i.RegisterToken(tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("RegisterToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_inMemory_RemoveUser(t *testing.T) {
	type fields struct {
		userById    map[model.UserIdentifier]model.User
		userByEmail map[string]model.User
		validToken  map[string]struct{}
	}
	type args struct {
		id model.UserIdentifier
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inMemory{
				userById:    tt.fields.userById,
				userByEmail: tt.fields.userByEmail,
				validToken:  tt.fields.validToken,
			}
			if err := i.RemoveUser(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("RemoveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_inMemory_SetUser(t *testing.T) {
	type fields struct {
		userById    map[model.UserIdentifier]model.User
		userByEmail map[string]model.User
		validToken  map[string]struct{}
	}
	type args struct {
		user model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inMemory{
				userById:    tt.fields.userById,
				userByEmail: tt.fields.userByEmail,
				validToken:  tt.fields.validToken,
			}
			if err := i.SetUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("SetUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_inMemory_UnregisterToken(t *testing.T) {
	type fields struct {
		userById    map[model.UserIdentifier]model.User
		userByEmail map[string]model.User
		validToken  map[string]struct{}
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inMemory{
				userById:    tt.fields.userById,
				userByEmail: tt.fields.userByEmail,
				validToken:  tt.fields.validToken,
			}
			if err := i.UnregisterToken(tt.args.token); (err != nil) != tt.wantErr {
				t.Errorf("UnregisterToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
