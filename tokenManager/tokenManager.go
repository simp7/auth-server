package tokenManager

import "auth-server/model"

type TokenManager interface {
	Generate(model.User) (string, error)
	Verify(accessToken string)
}
