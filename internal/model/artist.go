package model

type Artist struct {
	Base
	Name        string       `json:"username,omitempty"`
	Projects    []Project    `json:"projects,omitempty"`
	Withdrawals []Withdrawal `json:"withdrawals,omitempty"`
}
