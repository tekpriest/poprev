package model

type Deposit struct {
	Base
	Amount    float64 `json:"amount,omitempty"`
	UserID    string  `json:"user_id,omitempty"`
	ProjectID string  `json:"project_id,omitempty"`
}
