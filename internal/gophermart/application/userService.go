package application

import (
	"context"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/psfpro/gophermart/internal/gophermart/domain"
)

var ErrUserUnauthorized = errors.New("user unauthorized")

type UserService struct {
	userRepository domain.UserRepository
	hashService    AuthenticationService
}

func NewUserService(userRepository domain.UserRepository, hashService AuthenticationService) *UserService {
	return &UserService{userRepository: userRepository, hashService: hashService}
}

func (s *UserService) Registration(ctx context.Context, login string, password string) (*LoginResult, error) {
	UUID, err := uuid.NewV6()
	if err != nil {
		return nil, err
	}
	userID := domain.NewUserID(UUID)
	userLogin := domain.Login(login)
	passwordHash, err := s.hashService.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := domain.NewUser(
		userID,
		userLogin,
		passwordHash,
	)
	if err := s.userRepository.Save(ctx, user); err != nil {
		return nil, err
	}

	accessToken, err := s.hashService.AccessToken(user.ID().UUID)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.hashService.RefreshToken()
	if err != nil {
		return nil, err
	}

	return NewLoginResult(accessToken, refreshToken), nil
}

func (s *UserService) Login(ctx context.Context, login string, password string) (*LoginResult, error) {
	user, err := s.userRepository.GetByLogin(ctx, login)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, ErrUserUnauthorized
		}
		return nil, err
	}
	if err := s.hashService.CheckPassword(user.PasswordHash(), password); err != nil {
		return nil, err
	}

	accessToken, err := s.hashService.AccessToken(user.ID().UUID)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.hashService.RefreshToken()
	if err != nil {
		return nil, err
	}

	return NewLoginResult(accessToken, refreshToken), nil
}

type LoginResult struct {
	AccessToken  string
	RefreshToken string
}

func NewLoginResult(accessToken string, refreshToken string) *LoginResult {
	return &LoginResult{AccessToken: accessToken, RefreshToken: refreshToken}
}
