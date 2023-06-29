package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	Base
	Amount    float64 `json:"amount,omitempty"`
	Quantity  int     `json:"quantity,omitempty"`
	ProjectID string  `json:"project_id,omitempty"`
	UserID    string  `json:"user_id,omitempty"`
	Sales     []Sale  `json:"sales,omitempty"`
}

func (t *Token) BeforeCreate(d *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	t.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().String())

	return
}
