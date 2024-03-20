package domain

import (
	"context"
	"errors"
	"github.com/gofrs/uuid"
)

var ErrUserNotFound = errors.New("user not found")

//go:generate mockery --name UserRepository --with-expecter
type UserRepository interface {
	GetByID(ctx context.Context, userID uuid.UUID) (*User, error)
	GetByLogin(ctx context.Context, login string) (*User, error)
	Save(ctx context.Context, user *User) error
}
