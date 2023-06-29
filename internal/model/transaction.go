package model

import "github.com/google/uuid"

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
	EventID *uuid.UUID           `json:"event_id,omitempty"`
	Type    TransactionEventType `json:"event_type,omitempty"`
	Status  TransactionStatus    `json:"status,omitempty"`
}
