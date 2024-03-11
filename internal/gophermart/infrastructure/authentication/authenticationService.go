package authentication

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"github.com/psfpro/gophermart/internal/gophermart/domain"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const TokenExp = time.Hour * 3
const SecretKey = "supersecretkey"

type Claims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) HashPassword(password string) (domain.PasswordHash, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return domain.PasswordHash(bytes), nil
}

func (s *Service) CheckPassword(passwordHash domain.PasswordHash, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AccessToken(userID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExp)),
		},
		UserID: userID,
	})
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *Service) RefreshToken() (string, error) {
	bytes := make([]byte, 64)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (s *Service) GetUserID(tokenString string) (uuid.UUID, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	return claims.UserID, nil
}
