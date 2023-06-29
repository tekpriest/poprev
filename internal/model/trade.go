package model

type Trade struct {
	Base
	Quantity int     `json:"quantity,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
	UserID   string  `json:"user_id,omitempty"`
	SaleID   string  `json:"sale_id,omitempty"`
}
