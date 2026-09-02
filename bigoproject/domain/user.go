package domain

import (
	"errors"
	"strings"
)

var (
	ErrInvalidEmail      = errors.New("invalid email address")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrEserAlreadyExists = ErrUserAlreadyExists // Alias for typo compatibility
)

type User struct {
	id    string
	email string
	name  string
}

func NewUser(email, name string) (*User, error) {
	if strings.TrimSpace(email) == "" || !strings.Contains(email, "@") {
		return nil, ErrInvalidEmail
	}
	return &User{
		email: email,
		name:  name,
	}, nil
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Name() string {
	return u.name
}
