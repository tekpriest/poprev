package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/tekpriest/poprev/internal/constants"
)

type ProjectStatus string

const (
	ProjectPending  ProjectStatus = "pending"
	ProjectApproved ProjectStatus = "approved"
	ProjectDeclined ProjectStatus = "declined"
	ProjectComplete ProjectStatus = "complete"
	ProjectFlagged  ProjectStatus = "flagged"
	ProjectDisabled ProjectStatus = "inactive"
)

type Project struct {
	Base
	Amount      float64       `json:"amount,omitempty"`
	Tokens      int           `json:"tokens,omitempty"`
	Claimed     []Token       `json:"claimed,omitempty"     gorm:"foreignKey:UserID"`
	Status      ProjectStatus `json:"status,omitempty"`
	ArtistID    string        `json:"artist_id,omitempty"`
	Rate        Rate          `json:"rate,omitempty"`
	Withdrawals []Withdrawal  `json:"withdrawals,omitempty" gorm:"foreignKey:ProjectID"`
	Deposit     []Deposit     `json:"deposit,omitempty"`
}

func (p *Project) BeforeCreate(d *gorm.DB) (err error) {
	p.ID = uuid.New().String()
	p.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().String())
	p.Status = ProjectStatus(ProjectPending)

	return
}

func (p *Project) AfterCreate(d *gorm.DB) (err error) {
	if err := d.Create(&Rate{
		Buy:       constants.BUY_RATE,
		Sell:      constants.SELL_RATE,
		ProjectID: p.ID,
	}).Error; err != nil {
		return err
	}

	// automatically set project to approved
	d.Model(p).Update("status", ProjectStatus(ProjectApproved))

	return
}
