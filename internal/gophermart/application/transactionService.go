package application

import (
	"context"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/psfpro/gophermart/internal/gophermart/domain"
	"time"
)

type TransactionService struct {
	transactionRepository domain.TransactionRepository
}

func NewTransactionService(transactionRepository domain.TransactionRepository) *TransactionService {
	return &TransactionService{transactionRepository: transactionRepository}
}

var ErrTransactionUploadedByUser = errors.New("oldWithdrawal uploaded by user")
var ErrTransactionUploadedByOtherUser = errors.New("oldWithdrawal uploaded by other user")

func (s *TransactionService) NewAccrual(
	ctx context.Context,
	transactionID domain.TransactionID,
	userID domain.UserID,
	orderNumber domain.OrderNumber,
	now time.Time,
) error {
	transaction, err := s.transactionRepository.GetByOrderNumber(ctx, orderNumber)
	if err != nil && !errors.Is(err, domain.ErrTransactionNotFound) {
		return err
	}
	if transaction != nil && transaction.UserID() == userID {
		return ErrTransactionUploadedByUser
	} else if transaction != nil && transaction.UserID() != userID {
		return ErrTransactionUploadedByOtherUser
	}

	transaction = domain.NewTransaction(
		transactionID,
		userID,
		orderNumber,
		domain.TransactionStatusNew,
		domain.TransactionTypeAccrual,
		domain.TransactionAmount(0),
		now,
		now,
	)
	if err := s.transactionRepository.Save(ctx, transaction); err != nil {
		return err
	}

	return nil
}

func (s *TransactionService) GetAccruals(ctx context.Context, userID uuid.UUID) ([]*domain.Transaction, error) {
	return s.transactionRepository.GetAccrualsByUserID(ctx, userID)
}

func (s *TransactionService) NewWithdrawal(
	ctx context.Context,
	transactionID domain.TransactionID,
	userID domain.UserID,
	orderNumber domain.OrderNumber,
	amount domain.TransactionAmount,
	now time.Time,
) error {
	transaction, err := s.transactionRepository.GetByOrderNumber(ctx, orderNumber)
	if err != nil && !errors.Is(err, domain.ErrTransactionNotFound) {
		return err
	}
	if transaction != nil && transaction.UserID() == userID {
		return ErrTransactionUploadedByUser
	} else if transaction != nil && transaction.UserID() != userID {
		return ErrTransactionUploadedByOtherUser
	}

	transaction = domain.NewTransaction(
		transactionID,
		userID,
		orderNumber,
		domain.TransactionStatusProcessed,
		domain.TransactionTypeWithdrawal,
		amount,
		now,
		now,
	)
	if err := s.transactionRepository.Save(ctx, transaction); err != nil {
		return err
	}

	return nil
}

func (s *TransactionService) GetWithdrawals(ctx context.Context, userID uuid.UUID) ([]*domain.Transaction, error) {
	return s.transactionRepository.GetWithdrawalsByUserID(ctx, userID)
}

func (s *TransactionService) GetBalance(ctx context.Context, userID uuid.UUID) (*Account, error) {
	accruals, err := s.GetAccruals(ctx, userID)
	if err != nil {
		return nil, err
	}
	withdrawals, err := s.GetWithdrawals(ctx, userID)
	if err != nil {
		return nil, err
	}
	var accrualsSum int64
	var withdrawalsSum int64

	for _, v := range accruals {
		accrualsSum += int64(v.Amount())
	}
	for _, v := range withdrawals {
		withdrawalsSum += int64(v.Amount())
	}
	balance := accrualsSum - withdrawalsSum

	return &Account{
		Accruals:    domain.TransactionAmount(accrualsSum),
		Withdrawals: domain.TransactionAmount(withdrawalsSum),
		Balance:     domain.TransactionAmount(balance),
	}, nil
}

type Account struct {
	Accruals    domain.TransactionAmount
	Withdrawals domain.TransactionAmount
	Balance     domain.TransactionAmount
}
