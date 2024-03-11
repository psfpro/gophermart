package domain

import (
	"github.com/gofrs/uuid"
	"time"
)

const (
	TransactionStatusNew        TransactionStatus = "NEW"        // Заказ загружен в систему, но не попал в обработку
	TransactionStatusProcessing TransactionStatus = "PROCESSING" // Вознаграждение за заказ рассчитывается
	TransactionStatusInvalid    TransactionStatus = "INVALID"    // Система расчёта вознаграждений отказала в расчёте
	TransactionStatusProcessed  TransactionStatus = "PROCESSED"  // Данные по заказу проверены и информация о расчёте успешно получена
)

const (
	TransactionTypeAccrual    TransactionType = "Accrual"
	TransactionTypeWithdrawal TransactionType = "Withdrawal"
)

type TransactionID struct {
	uuid.UUID
}

func NewTransactionID(UUID uuid.UUID) TransactionID {
	return TransactionID{UUID: UUID}
}

type OrderNumber string

type TransactionStatus string

type TransactionAmount int64

func NewTransactionAmount(value float32) TransactionAmount {
	return TransactionAmount(value * 100)
}

func (a TransactionAmount) Value() float32 {
	return float32(a) / 100
}

type TransactionType string

type Transaction struct {
	id              TransactionID
	userID          UserID
	orderNumber     OrderNumber
	status          TransactionStatus
	transactionType TransactionType
	amount          TransactionAmount
	createdAt       time.Time
	updatedAt       time.Time
}

func (t *Transaction) ID() TransactionID {
	return t.id
}

func (t *Transaction) UserID() UserID {
	return t.userID
}

func (t *Transaction) OrderNumber() OrderNumber {
	return t.orderNumber
}

func (t *Transaction) Status() TransactionStatus {
	return t.status
}

func (t *Transaction) TransactionType() TransactionType {
	return t.transactionType
}

func (t *Transaction) Amount() TransactionAmount {
	return t.amount
}

func (t *Transaction) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Transaction) UpdatedAt() time.Time {
	return t.updatedAt
}

func NewTransaction(
	id TransactionID,
	userID UserID,
	orderNumber OrderNumber,
	status TransactionStatus,
	transactionType TransactionType,
	amount TransactionAmount,
	createdAt time.Time,
	updatedAt time.Time,
) *Transaction {
	return &Transaction{
		id:              id,
		userID:          userID,
		orderNumber:     orderNumber,
		status:          status,
		transactionType: transactionType,
		amount:          amount,
		createdAt:       createdAt,
		updatedAt:       updatedAt,
	}
}
