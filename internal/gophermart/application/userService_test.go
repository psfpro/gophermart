package application

import (
	"context"
	"github.com/gofrs/uuid"
	mocks2 "github.com/psfpro/gophermart/internal/gophermart/application/mocks"
	"github.com/psfpro/gophermart/internal/gophermart/domain"
	"github.com/psfpro/gophermart/internal/gophermart/domain/mocks"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewLoginResult(t *testing.T) {
	type args struct {
		accessToken  string
		refreshToken string
	}
	tests := []struct {
		name string
		args args
		want *LoginResult
	}{
		{
			name: "positive test",
			args: args{
				accessToken:  "",
				refreshToken: "",
			},
			want: &LoginResult{
				AccessToken:  "",
				RefreshToken: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoginResult(tt.args.accessToken, tt.args.refreshToken); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserService(t *testing.T) {
	userRepositoryMock := mocks.NewUserRepository(t)
	hashServiceMock := mocks2.NewAuthenticationService(t)
	type args struct {
		userRepository domain.UserRepository
		hashService    AuthenticationService
	}
	tests := []struct {
		name string
		args args
		want *UserService
	}{
		{
			name: "positive test",
			args: args{
				userRepository: userRepositoryMock,
				hashService:    hashServiceMock,
			},
			want: &UserService{
				userRepository: userRepositoryMock,
				hashService:    hashServiceMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.userRepository, tt.args.hashService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Login(t *testing.T) {
	type args struct {
		ctx      context.Context
		login    string
		password string
	}
	type expect struct {
		user             *domain.User
		err              error
		checkPasswordErr error
		userUUID         uuid.UUID
		accessToken      string
		refreshToken     string
	}
	tests := []struct {
		name    string
		args    args
		expect  expect
		want    *LoginResult
		wantErr error
	}{
		{
			name: "positive test",
			args: args{
				ctx:      nil,
				login:    "",
				password: "",
			},
			expect: expect{
				user: domain.NewUser(
					domain.NewUserID(uuid.UUID{}),
					"",
					"",
				),
				err:              nil,
				checkPasswordErr: nil,
				userUUID:         uuid.UUID{},
				accessToken:      "",
				refreshToken:     "",
			},
			want: &LoginResult{
				AccessToken:  "",
				RefreshToken: "",
			},
			wantErr: nil,
		},
		{
			name: "err user unauthorized",
			args: args{
				ctx:      nil,
				login:    "",
				password: "",
			},
			expect: expect{
				user:             nil,
				err:              domain.ErrUserNotFound,
				checkPasswordErr: nil,
				accessToken:      "",
				refreshToken:     "",
			},
			want: &LoginResult{
				AccessToken:  "",
				RefreshToken: "",
			},
			wantErr: ErrUserUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepositoryMock := mocks.NewUserRepository(t)
			userRepositoryMock.EXPECT().GetByLogin(tt.args.ctx, tt.args.login).Return(tt.expect.user, tt.expect.err).Maybe()
			hashServiceMock := mocks2.NewAuthenticationService(t)
			hashServiceMock.EXPECT().CheckPassword(domain.PasswordHash(tt.args.password), tt.args.password).Return(tt.expect.checkPasswordErr).Maybe()
			hashServiceMock.EXPECT().AccessToken(tt.expect.userUUID).Return(tt.expect.accessToken, nil).Maybe()
			hashServiceMock.EXPECT().RefreshToken().Return(tt.expect.refreshToken, nil).Maybe()
			s := &UserService{
				userRepository: userRepositoryMock,
				hashService:    hashServiceMock,
			}

			got, err := s.Login(tt.args.ctx, tt.args.login, tt.args.password)

			assert.ErrorIs(t, err, tt.wantErr)
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestUserService_Registration(t *testing.T) {
	type args struct {
		ctx      context.Context
		userID   domain.UserID
		login    string
		password string
	}
	type expect struct {
		oldUser      *domain.User
		user         *domain.User
		err          error
		userUUID     uuid.UUID
		accessToken  string
		refreshToken string
	}
	tests := []struct {
		name    string
		args    args
		expect  expect
		want    *LoginResult
		wantErr error
	}{
		{
			name: "positive test",
			args: args{
				ctx: nil,
				userID: domain.UserID{
					UUID: uuid.UUID{},
				},
				login:    "",
				password: "",
			},
			expect: expect{
				oldUser: nil,
				user: domain.NewUser(
					domain.NewUserID(uuid.UUID{}),
					"",
					"",
				),
				err:          nil,
				userUUID:     uuid.UUID{},
				accessToken:  "",
				refreshToken: "",
			},
			want: &LoginResult{
				AccessToken:  "",
				RefreshToken: "",
			},
			wantErr: nil,
		},
		{
			name: "err user login already taken",
			args: args{
				ctx: nil,
				userID: domain.UserID{
					UUID: uuid.UUID{},
				},
				login:    "",
				password: "",
			},
			expect: expect{
				oldUser: domain.NewUser(
					domain.NewUserID(uuid.UUID{}),
					"",
					"",
				),
				user:         nil,
				err:          nil,
				userUUID:     uuid.UUID{},
				accessToken:  "",
				refreshToken: "",
			},
			want: &LoginResult{
				AccessToken:  "",
				RefreshToken: "",
			},
			wantErr: ErrUserLoginAlreadyTaken,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userRepositoryMock := mocks.NewUserRepository(t)
			userRepositoryMock.EXPECT().GetByLogin(tt.args.ctx, tt.args.login).Return(tt.expect.oldUser, tt.expect.err).Maybe()
			userRepositoryMock.EXPECT().Save(tt.args.ctx, tt.expect.user).Return(nil).Maybe()
			hashServiceMock := mocks2.NewAuthenticationService(t)
			hashServiceMock.EXPECT().HashPassword(tt.args.password).Return(domain.PasswordHash(tt.args.password), nil).Maybe()
			hashServiceMock.EXPECT().AccessToken(tt.expect.userUUID).Return(tt.expect.accessToken, nil).Maybe()
			hashServiceMock.EXPECT().RefreshToken().Return(tt.expect.refreshToken, nil).Maybe()
			s := &UserService{
				userRepository: userRepositoryMock,
				hashService:    hashServiceMock,
			}

			got, err := s.Registration(tt.args.ctx, tt.args.userID, tt.args.login, tt.args.password)

			assert.ErrorIs(t, err, tt.wantErr)
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}
