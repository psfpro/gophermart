package domain

import (
	"github.com/gofrs/uuid"
	"reflect"
	"testing"
	"time"
)

func TestNewTransaction(t *testing.T) {
	type args struct {
		id              TransactionID
		userID          UserID
		orderNumber     OrderNumber
		status          TransactionStatus
		transactionType TransactionType
		amount          TransactionAmount
		createdAt       time.Time
		updatedAt       time.Time
	}
	tests := []struct {
		name string
		args args
		want *Transaction
	}{
		{
			name: "positive test",
			args: args{
				id: TransactionID{
					UUID: uuid.UUID{},
				},
				userID: UserID{
					UUID: uuid.UUID{},
				},
				orderNumber:     "",
				status:          "",
				transactionType: "",
				amount:          0,
				createdAt:       time.Time{},
				updatedAt:       time.Time{},
			},
			want: &Transaction{
				id: TransactionID{
					UUID: uuid.UUID{},
				},
				userID: UserID{
					UUID: uuid.UUID{},
				},
				orderNumber:     "",
				status:          "",
				transactionType: "",
				amount:          0,
				createdAt:       time.Time{},
				updatedAt:       time.Time{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransaction(tt.args.id, tt.args.userID, tt.args.orderNumber, tt.args.status, tt.args.transactionType, tt.args.amount, tt.args.createdAt, tt.args.updatedAt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTransactionAmount(t *testing.T) {
	type args struct {
		value float32
	}
	tests := []struct {
		name string
		args args
		want TransactionAmount
	}{
		{
			name: "positive test",
			args: args{
				value: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransactionAmount(tt.args.value); got != tt.want {
				t.Errorf("NewTransactionAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTransactionID(t *testing.T) {
	type args struct {
		UUID uuid.UUID
	}
	tests := []struct {
		name string
		args args
		want TransactionID
	}{
		{
			name: "positive test",
			args: args{
				UUID: uuid.UUID{},
			},
			want: TransactionID{
				UUID: uuid.UUID{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransactionID(tt.args.UUID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransactionID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionAmount_Value(t *testing.T) {
	tests := []struct {
		name string
		a    TransactionAmount
		want float32
	}{
		{
			name: "positive test",
			a:    0,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Value(); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_Amount(t1 *testing.T) {
	type fields struct {
		id              TransactionID
		userID          UserID
		orderNumber     OrderNumber
		status          TransactionStatus
		transactionType TransactionType
		amount          TransactionAmount
		createdAt       time.Time
		updatedAt       time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   TransactionAmount
	}{
		{
			name: "positive test",
			fields: fields{
				id: TransactionID{
					UUID: uuid.UUID{},
				},
				userID: UserID{
					UUID: uuid.UUID{},
				},
				orderNumber:     "",
				status:          "",
				transactionType: "",
				amount:          0,
				createdAt:       time.Time{},
				updatedAt:       time.Time{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transaction{
				id:              tt.fields.id,
				userID:          tt.fields.userID,
				orderNumber:     tt.fields.orderNumber,
				status:          tt.fields.status,
				transactionType: tt.fields.transactionType,
				amount:          tt.fields.amount,
				createdAt:       tt.fields.createdAt,
				updatedAt:       tt.fields.updatedAt,
			}
			if got := t.Amount(); got != tt.want {
				t1.Errorf("Amount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_CreatedAt(t1 *testing.T) {
	type fields struct {
		id              TransactionID
		userID          UserID
		orderNumber     OrderNumber
		status          TransactionStatus
		transactionType TransactionType
		amount          TransactionAmount
		createdAt       time.Time
		updatedAt       time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "positive test",
			fields: fields{
				id: TransactionID{
					UUID: uuid.UUID{},
				},
				userID: UserID{
					UUID: uuid.UUID{},
				},
				orderNumber:     "",
				status:          "",
				transactionType: "",
				amount:          0,
				createdAt:       time.Time{},
				updatedAt:       time.Time{},
			},
			want: time.Time{},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transaction{
				id:              tt.fields.id,
				userID:          tt.fields.userID,
				orderNumber:     tt.fields.orderNumber,
				status:          tt.fields.status,
				transactionType: tt.fields.transactionType,
				amount:          tt.fields.amount,
				createdAt:       tt.fields.createdAt,
				updatedAt:       tt.fields.updatedAt,
			}
			if got := t.CreatedAt(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("CreatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_ID(t1 *testing.T) {
	type fields struct {
		id              TransactionID
		userID          UserID
		orderNumber     OrderNumber
		status          TransactionStatus
		transactionType TransactionType
		amount          TransactionAmount
		createdAt       time.Time
		updatedAt       time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   TransactionID
	}{
		{
			name: "positive test",
			fields: fields{
				id: TransactionID{
					UUID: uuid.UUID{},
				},
				userID: UserID{
					UUID: uuid.UUID{},
				},
				orderNumber:     "",
				status:          "",
				transactionType: "",
				amount:          0,
				createdAt:       time.Time{},
				updatedAt:       time.Time{},
			},
			want: TransactionID{
				UUID: uuid.UUID{},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transaction{
				id:              tt.fields.id,
				userID:          tt.fields.userID,
				orderNumber:     tt.fields.orderNumber,
				status:          tt.fields.status,
				transactionType: tt.fields.transactionType,
				amount:          tt.fields.amount,
				createdAt:       tt.fields.createdAt,
				updatedAt:       tt.fields.updatedAt,
			}
			if got := t.ID(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_OrderNumber(t1 *testing.T) {
	type fields struct {
		id              TransactionID
		userID          UserID
		orderNumber     OrderNumber
		status          TransactionStatus
		transactionType TransactionType
		amount          TransactionAmount
		createdAt       time.Time
		updatedAt       time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   OrderNumber
	}{
		{
			name: "positive test",
			fields: fields{
				id: TransactionID{
					UUID: uuid.UUID{},
				},
				userID: UserID{
					UUID: uuid.UUID{},
				},
				orderNumber:     "",
				status:          "",
				transactionType: "",
				amount:          0,
				createdAt:       time.Time{},
				updatedAt:       time.Time{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transaction{
				id:              tt.fields.id,
				userID:          tt.fields.userID,
				orderNumber:     tt.fields.orderNumber,
				status:          tt.fields.status,
				transactionType: tt.fields.transactionType,
				amount:          tt.fields.amount,
				createdAt:       tt.fields.createdAt,
				updatedAt:       tt.fields.updatedAt,
			}
			if got := t.OrderNumber(); got != tt.want {
				t1.Errorf("OrderNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_Processed(t1 *testing.T) {
	type fields struct {
		id              TransactionID
		userID          UserID
		orderNumber     OrderNumber
		status          TransactionStatus
		transactionType TransactionType
		amount          TransactionAmount
		createdAt       time.Time
		updatedAt       time.Time
	}
	type args struct {
		status TransactionStatus
		amount TransactionAmount
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "positive test",
			fields: fields{
				id: TransactionID{
					UUID: uuid.UUID{},
				},
				userID: UserID{
					UUID: uuid.UUID{},
				},
				orderNumber:     "",
				status:          "",
				transactionType: "",
				amount:          0,
				createdAt:       time.Time{},
				updatedAt:       time.Time{},
			},
			args: args{
				status: "",
				amount: 0,
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transaction{
				id:              tt.fields.id,
				userID:          tt.fields.userID,
				orderNumber:     tt.fields.orderNumber,
				status:          tt.fields.status,
				transactionType: tt.fields.transactionType,
				amount:          tt.fields.amount,
				createdAt:       tt.fields.createdAt,
				updatedAt:       tt.fields.updatedAt,
			}
			t.Processed(tt.args.status, tt.args.amount)
		})
	}
}

func TestTransaction_Status(t1 *testing.T) {
	type fields struct {
		id              TransactionID
		userID          UserID
		orderNumber     OrderNumber
		status          TransactionStatus
		transactionType TransactionType
		amount          TransactionAmount
		createdAt       time.Time
		updatedAt       time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   TransactionStatus
	}{
		{
			name: "positive test",
			fields: fields{
				id: TransactionID{
					UUID: uuid.UUID{},
				},
				userID: UserID{
					UUID: uuid.UUID{},
				},
				orderNumber:     "",
				status:          "",
				transactionType: "",
				amount:          0,
				createdAt:       time.Time{},
				updatedAt:       time.Time{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transaction{
				id:              tt.fields.id,
				userID:          tt.fields.userID,
				orderNumber:     tt.fields.orderNumber,
				status:          tt.fields.status,
				transactionType: tt.fields.transactionType,
				amount:          tt.fields.amount,
				createdAt:       tt.fields.createdAt,
				updatedAt:       tt.fields.updatedAt,
			}
			if got := t.Status(); got != tt.want {
				t1.Errorf("Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_TransactionType(t1 *testing.T) {
	type fields struct {
		id              TransactionID
		userID          UserID
		orderNumber     OrderNumber
		status          TransactionStatus
		transactionType TransactionType
		amount          TransactionAmount
		createdAt       time.Time
		updatedAt       time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   TransactionType
	}{
		{
			name: "positive test",
			fields: fields{
				id: TransactionID{
					UUID: uuid.UUID{},
				},
				userID: UserID{
					UUID: uuid.UUID{},
				},
				orderNumber:     "",
				status:          "",
				transactionType: "",
				amount:          0,
				createdAt:       time.Time{},
				updatedAt:       time.Time{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transaction{
				id:              tt.fields.id,
				userID:          tt.fields.userID,
				orderNumber:     tt.fields.orderNumber,
				status:          tt.fields.status,
				transactionType: tt.fields.transactionType,
				amount:          tt.fields.amount,
				createdAt:       tt.fields.createdAt,
				updatedAt:       tt.fields.updatedAt,
			}
			if got := t.TransactionType(); got != tt.want {
				t1.Errorf("TransactionType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_UpdatedAt(t1 *testing.T) {
	type fields struct {
		id              TransactionID
		userID          UserID
		orderNumber     OrderNumber
		status          TransactionStatus
		transactionType TransactionType
		amount          TransactionAmount
		createdAt       time.Time
		updatedAt       time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "positive test",
			fields: fields{
				id: TransactionID{
					UUID: uuid.UUID{},
				},
				userID: UserID{
					UUID: uuid.UUID{},
				},
				orderNumber:     "",
				status:          "",
				transactionType: "",
				amount:          0,
				createdAt:       time.Time{},
				updatedAt:       time.Time{},
			},
			want: time.Time{},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transaction{
				id:              tt.fields.id,
				userID:          tt.fields.userID,
				orderNumber:     tt.fields.orderNumber,
				status:          tt.fields.status,
				transactionType: tt.fields.transactionType,
				amount:          tt.fields.amount,
				createdAt:       tt.fields.createdAt,
				updatedAt:       tt.fields.updatedAt,
			}
			if got := t.UpdatedAt(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("UpdatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_UserID(t1 *testing.T) {
	type fields struct {
		id              TransactionID
		userID          UserID
		orderNumber     OrderNumber
		status          TransactionStatus
		transactionType TransactionType
		amount          TransactionAmount
		createdAt       time.Time
		updatedAt       time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   UserID
	}{
		{
			name: "positive test",
			fields: fields{
				id: TransactionID{
					UUID: uuid.UUID{},
				},
				userID: UserID{
					UUID: uuid.UUID{},
				},
				orderNumber:     "",
				status:          "",
				transactionType: "",
				amount:          0,
				createdAt:       time.Time{},
				updatedAt:       time.Time{},
			},
			want: UserID{
				UUID: uuid.UUID{},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Transaction{
				id:              tt.fields.id,
				userID:          tt.fields.userID,
				orderNumber:     tt.fields.orderNumber,
				status:          tt.fields.status,
				transactionType: tt.fields.transactionType,
				amount:          tt.fields.amount,
				createdAt:       tt.fields.createdAt,
				updatedAt:       tt.fields.updatedAt,
			}
			if got := t.UserID(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("UserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
