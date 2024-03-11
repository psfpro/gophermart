package postgres

import (
	"github.com/gofrs/uuid"
	"github.com/psfpro/gophermart/internal/gophermart/domain"
	"time"
)

type TransactionData struct {
	id              uuid.UUID
	userID          uuid.UUID
	orderNumber     string
	status          string
	transactionType string
	amount          int64
	createdAt       time.Time
	updatedAt       time.Time
}

func NewTransactionDataFromEntity(transaction *domain.Transaction) *TransactionData {
	return &TransactionData{
		id:              transaction.ID().UUID,
		userID:          transaction.UserID().UUID,
		orderNumber:     string(transaction.OrderNumber()),
		status:          string(transaction.Status()),
		transactionType: string(transaction.TransactionType()),
		amount:          int64(transaction.Amount()),
		createdAt:       transaction.CreatedAt(),
		updatedAt:       transaction.UpdatedAt(),
	}
}

func (d TransactionData) entity() (*domain.Transaction, error) {
	return domain.NewTransaction(
		domain.NewTransactionID(d.id),
		domain.NewUserID(d.userID),
		domain.OrderNumber(d.orderNumber),
		domain.TransactionStatus(d.status),
		domain.TransactionType(d.transactionType),
		domain.TransactionAmount(d.amount),
		d.createdAt,
		d.updatedAt,
	), nil
}
