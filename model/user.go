package model

import (
	"golang.org/x/crypto/bcrypt"
)

type UserIdentifier struct {
	Uid uint64
}

type User struct {
	UserIdentifier
	Email    string
	Password string
	Nickname string
	Role     []string
}

func (u User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
