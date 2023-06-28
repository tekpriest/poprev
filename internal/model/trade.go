package model

import "github.com/google/uuid"

type Trade struct {
	Base
	Quantity int        `json:"quantity,omitempty"`
	Amount   float64    `json:"amount,omitempty"`
	UserID   *uuid.UUID `json:"user_id,omitempty"`
	SaleID   *uuid.UUID `json:"sale_id,omitempty"`
}
