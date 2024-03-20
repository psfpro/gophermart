package domain

import (
	"github.com/gofrs/uuid"
)

type UserID struct {
	uuid.UUID
}

func NewUserID(UUID uuid.UUID) UserID {
	return UserID{UUID: UUID}
}

type Login string

type PasswordHash string

type User struct {
	id           UserID
	login        Login
	passwordHash PasswordHash
}

func (u *User) ID() UserID {
	return u.id
}

func (u *User) Login() Login {
	return u.login
}

func (u *User) PasswordHash() PasswordHash {
	return u.passwordHash
}

func NewUser(id UserID, login Login, passwordHash PasswordHash) *User {
	return &User{id: id, login: login, passwordHash: passwordHash}
}
