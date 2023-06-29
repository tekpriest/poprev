package model

type Token struct {
	Base
	Amount    float64 `json:"amount,omitempty"`
	Quantity  int     `json:"quantity,omitempty"`
	ProjectID string  `json:"project_id,omitempty"`
	UserID    string  `json:"user_id,omitempty"`
}
