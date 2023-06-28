package model

import "github.com/google/uuid"

type ProjectStatus string

type Project struct {
	Base
	Amount      float64       `json:"amount,omitempty"`
	Tokens      int           `json:"tokens,omitempty"`
	Claimed     []Token       `json:"claimed,omitempty"`
	Status      ProjectStatus `json:"status,omitempty"`
	ArtistID    *uuid.UUID    `json:"artist_id,omitempty"`
	Rate        Rate          `json:"rate,omitempty"`
	Withdrawals []Withdrawal  `json:"withdrawals,omitempty"`
	Deposit     []Deposit     `json:"deposit,omitempty"`
}
