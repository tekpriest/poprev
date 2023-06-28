package model

type User struct {
	Base
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"-,omitempty"`
	Email     string    `json:"email,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Tokens    []Token   `json:"tokens,omitempty"`
	Deposits  []Deposit `json:"deposits,omitempty"`
	Trades    []Trade   `json:"trades,omitempty"`
	Sales     []Sale    `json:"sales,omitempty"`
}
