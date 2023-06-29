package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/tekpriest/poprev/internal/constants"
)

type SaleStatus string

const (
	OnSale        SaleStatus = "on_sale"
	SaleCompleted SaleStatus = "completed"
	SaleCancelled SaleStatus = "cancelled"
)

type Sale struct {
	Base
	Quantity int        `json:"quantity,omitempty"`
	MinOrder int        `json:"min_order,omitempty"`
	MaxOrder int        `json:"max_order,omitempty"`
	Rate     float32    `json:"rate,omitempty"`
	Status   SaleStatus `json:"status,omitempty"`
	TokenID  string     `json:"token_id,omitempty"`
	SellerID string     `json:"seller_id,omitempty"`
	Trades   []Trade    `json:"trades,omitempty"`
}

func (s *Sale) BeforeCreate(d *gorm.DB) (err error) {
	s.ID = uuid.New().String()
	s.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().String())
	s.Status = SaleStatus(OnSale)
	s.Rate = constants.SELL_RATE

	return
}
