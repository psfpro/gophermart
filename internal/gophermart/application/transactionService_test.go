package application

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/psfpro/gophermart/internal/gophermart/domain"
	"github.com/psfpro/gophermart/internal/gophermart/domain/mocks"
	"reflect"
	"testing"
	"time"
)

func TestNewTransactionService(t *testing.T) {
	transactionRepositoryMock := mocks.NewTransactionRepository(t)
	type args struct {
		transactionRepository domain.TransactionRepository
	}
	tests := []struct {
		name string
		args args
		want *TransactionService
	}{
		{
			name: "test positive",
			args: args{
				transactionRepository: transactionRepositoryMock,
			},
			want: &TransactionService{
				transactionRepository: transactionRepositoryMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransactionService(tt.args.transactionRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransactionService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionService_GetAccruals(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	type expect struct {
		data []*domain.Transaction
		err  error
	}
	tests := []struct {
		name    string
		args    args
		expect  expect
		want    []*domain.Transaction
		wantErr bool
	}{
		{
			name: "positive test",
			args: args{
				ctx:    nil,
				userID: uuid.UUID{},
			},
			expect: expect{
				data: []*domain.Transaction{},
				err:  nil,
			},
			want:    []*domain.Transaction{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRepositoryMock := mocks.NewTransactionRepository(t)
			transactionRepositoryMock.EXPECT().GetAccrualsByUserID(tt.args.ctx, tt.args.userID).Return(tt.expect.data, tt.expect.err)
			s := &TransactionService{
				transactionRepository: transactionRepositoryMock,
			}
			got, err := s.GetAccruals(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccruals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccruals() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionService_GetBalance(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	type expect struct {
		accruals       []*domain.Transaction
		withdrawals    []*domain.Transaction
		accrualsErr    error
		withdrawalsErr error
	}
	tests := []struct {
		name    string
		args    args
		expect  expect
		want    *Account
		wantErr bool
	}{
		{
			name: "positive test",
			args: args{
				ctx:    nil,
				userID: uuid.UUID{},
			},
			expect: expect{
				accruals:       []*domain.Transaction{},
				withdrawals:    []*domain.Transaction{},
				accrualsErr:    nil,
				withdrawalsErr: nil,
			},
			want: &Account{
				Accruals:    0,
				Withdrawals: 0,
				Balance:     0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRepositoryMock := mocks.NewTransactionRepository(t)
			transactionRepositoryMock.EXPECT().GetAccrualsByUserID(tt.args.ctx, tt.args.userID).Return(tt.expect.accruals, tt.expect.accrualsErr)
			transactionRepositoryMock.EXPECT().GetWithdrawalsByUserID(tt.args.ctx, tt.args.userID).Return(tt.expect.withdrawals, tt.expect.withdrawalsErr)
			s := &TransactionService{
				transactionRepository: transactionRepositoryMock,
			}
			got, err := s.GetBalance(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBalance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionService_GetWithdrawals(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	type expect struct {
		withdrawals []*domain.Transaction
		err         error
	}
	tests := []struct {
		name    string
		args    args
		expect  expect
		want    []*domain.Transaction
		wantErr bool
	}{
		{
			name: "positive test",
			args: args{
				ctx:    nil,
				userID: uuid.UUID{},
			},
			expect: expect{
				withdrawals: []*domain.Transaction{},
				err:         nil,
			},
			want:    []*domain.Transaction{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRepositoryMock := mocks.NewTransactionRepository(t)
			transactionRepositoryMock.EXPECT().GetWithdrawalsByUserID(tt.args.ctx, tt.args.userID).Return(tt.expect.withdrawals, tt.expect.err)
			s := &TransactionService{
				transactionRepository: transactionRepositoryMock,
			}
			got, err := s.GetWithdrawals(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWithdrawals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWithdrawals() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionService_NewAccrual(t *testing.T) {
	type args struct {
		ctx           context.Context
		transactionID domain.TransactionID
		userID        domain.UserID
		orderNumber   domain.OrderNumber
		now           time.Time
	}
	type expect struct {
		oldAccrual *domain.Transaction
		err        error
		accrual    *domain.Transaction
	}
	tests := []struct {
		name    string
		args    args
		expect  expect
		wantErr bool
	}{
		{
			name: "positive test",
			args: args{
				ctx: nil,
				transactionID: domain.TransactionID{
					UUID: uuid.UUID{},
				},
				userID: domain.UserID{
					UUID: uuid.UUID{},
				},
				orderNumber: "",
				now:         time.Time{},
			},
			expect: expect{
				oldAccrual: nil,
				err:        nil,
				accrual: domain.NewTransaction(
					domain.NewTransactionID(uuid.UUID{}),
					domain.NewUserID(uuid.UUID{}),
					"",
					domain.TransactionStatusNew,
					domain.TransactionTypeAccrual,
					0,
					time.Time{},
					time.Time{},
				),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRepositoryMock := mocks.NewTransactionRepository(t)
			transactionRepositoryMock.EXPECT().GetByOrderNumber(tt.args.ctx, tt.args.orderNumber).Return(tt.expect.oldAccrual, tt.expect.err)
			transactionRepositoryMock.EXPECT().Save(tt.args.ctx, tt.expect.accrual).Return(nil)
			s := &TransactionService{
				transactionRepository: transactionRepositoryMock,
			}
			if err := s.NewAccrual(tt.args.ctx, tt.args.transactionID, tt.args.userID, tt.args.orderNumber, tt.args.now); (err != nil) != tt.wantErr {
				t.Errorf("NewAccrual() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransactionService_NewWithdrawal(t *testing.T) {
	type args struct {
		ctx           context.Context
		transactionID domain.TransactionID
		userID        domain.UserID
		orderNumber   domain.OrderNumber
		amount        domain.TransactionAmount
		now           time.Time
	}
	type expect struct {
		oldWithdrawal *domain.Transaction
		err           error
		withdrawal    *domain.Transaction
	}
	tests := []struct {
		name    string
		args    args
		expect  expect
		wantErr bool
	}{
		{
			name: "positive test",
			args: args{
				ctx: nil,
				userID: domain.UserID{
					UUID: uuid.UUID{},
				},
				orderNumber: "",
				amount:      0,
				now:         time.Time{},
			},
			expect: expect{
				oldWithdrawal: nil,
				err:           nil,
				withdrawal: domain.NewTransaction(
					domain.NewTransactionID(uuid.UUID{}),
					domain.NewUserID(uuid.UUID{}),
					"",
					domain.TransactionStatusProcessed,
					domain.TransactionTypeWithdrawal,
					0,
					time.Time{},
					time.Time{},
				),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRepositoryMock := mocks.NewTransactionRepository(t)
			transactionRepositoryMock.EXPECT().GetByOrderNumber(tt.args.ctx, tt.args.orderNumber).Return(tt.expect.oldWithdrawal, tt.expect.err)
			transactionRepositoryMock.EXPECT().Save(tt.args.ctx, tt.expect.withdrawal).Return(nil)
			s := &TransactionService{
				transactionRepository: transactionRepositoryMock,
			}
			if err := s.NewWithdrawal(tt.args.ctx, tt.args.transactionID, tt.args.userID, tt.args.orderNumber, tt.args.amount, tt.args.now); (err != nil) != tt.wantErr {
				t.Errorf("NewWithdrawal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
