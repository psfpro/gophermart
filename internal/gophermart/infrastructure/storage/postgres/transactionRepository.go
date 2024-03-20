package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/psfpro/gophermart/internal/gophermart/domain"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) CreateTable(ctx context.Context) error {
	query := `
CREATE TABLE IF NOT EXISTS "transaction" (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    order_number VARCHAR(255),
    status VARCHAR(16),
    transaction_type VARCHAR(16),
    amount BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
`
	_, err := r.db.ExecContext(ctx, query)

	return err
}

func (r *TransactionRepository) GetByOrderNumber(ctx context.Context, orderNumber domain.OrderNumber) (*domain.Transaction, error) {
	row := r.db.QueryRowContext(ctx, `
SELECT id, user_id, order_number, status, transaction_type, amount, created_at, updated_at
FROM "transaction"
WHERE order_number=$1
`, orderNumber)
	data := TransactionData{}
	err := row.Scan(
		&data.id,
		&data.userID,
		&data.orderNumber,
		&data.status,
		&data.transactionType,
		&data.amount,
		&data.createdAt,
		&data.updatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrTransactionNotFound
		}
		return nil, err
	}

	return data.entity()
}

func (r *TransactionRepository) GetAccrualsByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Transaction, error) {
	var res []*domain.Transaction
	rows, err := r.db.QueryContext(ctx, `
SELECT id, user_id, order_number, status, transaction_type, amount, created_at, updated_at
FROM "transaction"
WHERE user_id=$1 AND transaction_type=$2`, userID, domain.TransactionTypeAccrual)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		data := TransactionData{}
		err = rows.Scan(
			&data.id,
			&data.userID,
			&data.orderNumber,
			&data.status,
			&data.transactionType,
			&data.amount,
			&data.createdAt,
			&data.updatedAt,
		)
		if err != nil {
			return nil, err
		}
		entity, err := data.entity()
		if err != nil {
			return nil, err
		}
		res = append(res, entity)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *TransactionRepository) GetNewAccruals(ctx context.Context) ([]*domain.Transaction, error) {
	var res []*domain.Transaction
	rows, err := r.db.QueryContext(ctx, `
SELECT id, user_id, order_number, status, transaction_type, amount, created_at, updated_at
FROM "transaction"
WHERE status=$1 AND transaction_type=$2`, domain.TransactionStatusNew, domain.TransactionTypeAccrual)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		data := TransactionData{}
		err = rows.Scan(
			&data.id,
			&data.userID,
			&data.orderNumber,
			&data.status,
			&data.transactionType,
			&data.amount,
			&data.createdAt,
			&data.updatedAt,
		)
		if err != nil {
			return nil, err
		}
		entity, err := data.entity()
		if err != nil {
			return nil, err
		}
		res = append(res, entity)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *TransactionRepository) GetWithdrawalsByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Transaction, error) {
	var res []*domain.Transaction
	rows, err := r.db.QueryContext(ctx, `
SELECT id, user_id, order_number, status, transaction_type, amount, created_at, updated_at
FROM "transaction"
WHERE user_id=$1 AND transaction_type=$2`, userID, domain.TransactionTypeWithdrawal)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		data := TransactionData{}
		err = rows.Scan(
			&data.id,
			&data.userID,
			&data.orderNumber,
			&data.status,
			&data.transactionType,
			&data.amount,
			&data.createdAt,
			&data.updatedAt,
		)
		if err != nil {
			return nil, err
		}
		entity, err := data.entity()
		if err != nil {
			return nil, err
		}
		res = append(res, entity)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *TransactionRepository) Save(ctx context.Context, transaction *domain.Transaction) error {
	data := NewTransactionDataFromEntity(transaction)
	query := `
INSERT INTO "transaction" (id, user_id, order_number, status, transaction_type, amount, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (id)
DO UPDATE SET
    user_id = $2,
    order_number = $3,
    status = $4,
    transaction_type = $5,
    amount = $6,
    created_at = $7,
    updated_at = $8
`

	stm, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stm.Exec(
		data.id,
		data.userID,
		data.orderNumber,
		data.status,
		data.transactionType,
		data.amount,
		data.createdAt,
		data.updatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
