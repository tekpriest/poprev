package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Withdrawal struct {
	Base
	Amount    float64 `json:"amount,omitempty"`
	ArtistID  string  `json:"aritst_id,omitempty"`
	ProjectID string  `json:"project_id,omitempty"`
}

func (w *Withdrawal) BeforeCreate(d *gorm.DB) (err error) {
	w.ID = uuid.New().String()
	w.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().String())

	return
}
