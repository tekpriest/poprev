package model

import "github.com/google/uuid"

type Token struct {
	Base
	Amount    float64    `json:"amount,omitempty"`
	Quantity  int        `json:"quantity,omitempty"`
	ProjectID *uuid.UUID `json:"project_id,omitempty"`
	UserID    *uuid.UUID `json:"user_id,omitempty"`
}
