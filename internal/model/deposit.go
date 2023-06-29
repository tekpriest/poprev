package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Deposit struct {
	ID             string    `json:"id,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	Amount         float64   `json:"amount,omitempty"`
	UserID         string    `json:"user_id,omitempty"`
	TokenRequested int       `json:"tokens"               gorm:"column:tokens"`
	ProjectID      string    `json:"project_id,omitempty"`
}

func (de *Deposit) BeforeCreate(d *gorm.DB) (err error) {
	de.ID = uuid.New().String()
	de.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().String())

	return
}
