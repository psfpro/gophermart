package domain

import (
	"context"
	"errors"
	"github.com/gofrs/uuid"
)

var ErrTransactionNotFound = errors.New("transaction not found")

//go:generate mockery --name TransactionRepository --with-expecter
type TransactionRepository interface {
	GetByOrderNumber(ctx context.Context, orderNumber OrderNumber) (*Transaction, error)
	GetAccrualsByUserID(ctx context.Context, userID uuid.UUID) ([]*Transaction, error)
	GetNewAccruals(ctx context.Context) ([]*Transaction, error)
	GetWithdrawalsByUserID(ctx context.Context, userID uuid.UUID) ([]*Transaction, error)
	Save(ctx context.Context, transaction *Transaction) error
}
