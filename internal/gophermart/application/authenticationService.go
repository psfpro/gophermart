package application

import (
	"github.com/gofrs/uuid"
	"github.com/psfpro/gophermart/internal/gophermart/domain"
)

type AuthenticationService interface {
	HashPassword(password string) (domain.PasswordHash, error)
	CheckPassword(passwordHash domain.PasswordHash, providedPassword string) error
	AccessToken(userID uuid.UUID) (string, error)
	RefreshToken() (string, error)
	GetUserID(tokenString string) (uuid.UUID, error)
}
