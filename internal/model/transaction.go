package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	TransactionEventType string
	TransactionStatus    string
)

const (
	DepositEvent    TransactionEventType = "deposit"
	WithdrawalEvent TransactionEventType = "withdrawal"
	SaleEvent       TransactionEventType = "sale"
)

const (
	TransactionPending    TransactionStatus = "pending"
	TransactionFailed     TransactionStatus = "failed"
	TransactionSuccessful TransactionStatus = "successful"
)

type Transaction struct {
	Base
	Fee     float32              `json:"fee,omitempty"`
	Amount  float64              `json:"amount,omitempty"`
	EventID string               `json:"event_id,omitempty"`
	Type    TransactionEventType `json:"event_type,omitempty" gorm:"column:event_type"`
	Status  TransactionStatus    `json:"status,omitempty"`
}

func (t *Transaction) BeforeCreate(d *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	t.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().String())
	t.Status = TransactionStatus(TransactionPending)

	return
}
