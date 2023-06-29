package model

type Withdrawal struct {
	Base
	Amount    float64 `json:"amount,omitempty"`
	ArtistID  string  `json:"aritst_id,omitempty"`
	ProjectID string  `json:"project_id,omitempty"`
}
