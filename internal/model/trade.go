package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Trade struct {
	Base
	Quantity int     `json:"quantity,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
	UserID   string  `json:"user_id,omitempty"`
	SaleID   string  `json:"sale_id,omitempty"`
}

func (t *Trade) BeforeCreate(d *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	t.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().String())

	return
}
