package postgres

import (
	"github.com/gofrs/uuid"
	"github.com/psfpro/gophermart/internal/gophermart/domain"
)

type UserData struct {
	id           uuid.UUID
	login        string
	passwordHash string
}

func NewUserDataFromEntity(user *domain.User) *UserData {
	return &UserData{
		id:           user.ID().UUID,
		login:        string(user.Login()),
		passwordHash: string(user.PasswordHash()),
	}
}

func (m UserData) entity() (*domain.User, error) {
	return domain.NewUser(
		domain.NewUserID(m.id),
		domain.Login(m.login),
		domain.PasswordHash(m.passwordHash),
	), nil
}
