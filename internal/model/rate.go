package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rate struct {
	Base
	Buy       float32 `json:"buy,omitempty"`
	Sell      float32 `json:"sell,omitempty"`
	ProjectID string  `json:"project_id,omitempty"`
}

func (p *Rate) BeforeCreate(d *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	p.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().String())

	return
}
