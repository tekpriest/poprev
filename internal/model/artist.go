package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Artist struct {
	ID          string       `json:"id,omitempty"`
	Name        string       `json:"username,omitempty"`
	CreatedAt   time.Time    `json:"created_at,omitempty"`
	Projects    []Project    `json:"projects,omitempty"`
	Withdrawals []Withdrawal `json:"withdrawals,omitempty" gorm:"foreignKey:ArtistID"`
}

func (a *Artist) BeforeCreate(d *gorm.DB) (err error) {
	a.ID = uuid.New().String()
	a.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().String())

	return
}
