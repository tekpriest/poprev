package model

import "github.com/google/uuid"

type Withdrawal struct {
	Base
	Amount    float64    `json:"amount,omitempty"`
	AritstID  *uuid.UUID `json:"aritst_id,omitempty"`
	ProjectID *uuid.UUID `json:"project_id,omitempty"`
}
