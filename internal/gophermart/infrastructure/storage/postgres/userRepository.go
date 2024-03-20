package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/psfpro/gophermart/internal/gophermart/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateTable(ctx context.Context) error {
	query := `
CREATE TABLE IF NOT EXISTS "user" (
    id UUID PRIMARY KEY,
    login VARCHAR(255),
    password_hash VARCHAR(255)
);
`
	_, err := r.db.ExecContext(ctx, query)

	return err
}

func (r *UserRepository) GetByID(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, login, password_hash FROM "user" where id=$1`, userID)
	data := UserData{}
	err := row.Scan(&data.id, &data.login, &data.passwordHash)
	if err != nil {
		return nil, err
	}

	return data.entity()
}

func (r *UserRepository) GetByLogin(ctx context.Context, login string) (*domain.User, error) {
	row := r.db.QueryRowContext(ctx, `SELECT id, login, password_hash FROM "user" where login=$1`, login)
	data := UserData{}
	err := row.Scan(&data.id, &data.login, &data.passwordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	return data.entity()
}

func (r *UserRepository) Save(ctx context.Context, user *domain.User) error {
	data := NewUserDataFromEntity(user)
	query := `
INSERT INTO "user" (id, login, password_hash)
VALUES ($1, $2, $3)
ON CONFLICT (id)
DO UPDATE SET
    login = $2,
    password_hash = $3
`

	stm, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stm.Exec(data.id, data.login, data.passwordHash)
	if err != nil {
		return err
	}

	return nil
}
