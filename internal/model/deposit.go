package model

import "github.com/google/uuid"

type Deposit struct {
	Base
	Amount    float64    `json:"amount,omitempty"`
	UserID    *uuid.UUID `json:"user_id,omitempty"`
	ProjectID *uuid.UUID `json:"project_id,omitempty"`
}
