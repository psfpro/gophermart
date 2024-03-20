package domain

import (
	"github.com/gofrs/uuid"
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		id           UserID
		login        Login
		passwordHash PasswordHash
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		{
			name: "positive test",
			args: args{
				id: UserID{
					UUID: uuid.UUID{},
				},
				login:        "",
				passwordHash: "",
			},
			want: &User{
				id: UserID{
					UUID: uuid.UUID{},
				},
				login:        "",
				passwordHash: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.id, tt.args.login, tt.args.passwordHash); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserID(t *testing.T) {
	type args struct {
		UUID uuid.UUID
	}
	tests := []struct {
		name string
		args args
		want UserID
	}{
		{
			name: "positive test",
			args: args{
				UUID: uuid.UUID{},
			},
			want: UserID{
				UUID: uuid.UUID{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserID(tt.args.UUID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_ID(t *testing.T) {
	type fields struct {
		id           UserID
		login        Login
		passwordHash PasswordHash
	}
	tests := []struct {
		name   string
		fields fields
		want   UserID
	}{
		{
			name: "positive test",
			fields: fields{
				id: UserID{
					UUID: uuid.UUID{},
				},
				login:        "",
				passwordHash: "",
			},
			want: UserID{
				UUID: uuid.UUID{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				id:           tt.fields.id,
				login:        tt.fields.login,
				passwordHash: tt.fields.passwordHash,
			}
			if got := u.ID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Login(t *testing.T) {
	type fields struct {
		id           UserID
		login        Login
		passwordHash PasswordHash
	}
	tests := []struct {
		name   string
		fields fields
		want   Login
	}{
		{
			name: "positive test",
			fields: fields{
				id: UserID{
					UUID: uuid.UUID{},
				},
				login:        "",
				passwordHash: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				id:           tt.fields.id,
				login:        tt.fields.login,
				passwordHash: tt.fields.passwordHash,
			}
			if got := u.Login(); got != tt.want {
				t.Errorf("Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_PasswordHash(t *testing.T) {
	type fields struct {
		id           UserID
		login        Login
		passwordHash PasswordHash
	}
	tests := []struct {
		name   string
		fields fields
		want   PasswordHash
	}{
		{
			name: "positive test",
			fields: fields{
				id: UserID{
					UUID: uuid.UUID{},
				},
				login:        "",
				passwordHash: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				id:           tt.fields.id,
				login:        tt.fields.login,
				passwordHash: tt.fields.passwordHash,
			}
			if got := u.PasswordHash(); got != tt.want {
				t.Errorf("PasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
