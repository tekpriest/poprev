package user

import "github.com/gookit/validate"

type UpdateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type RegisterUser struct {
	Email     string `json:"email"      validate:"required"`
	Password  string `json:"password"   validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"   validate:"required"`
}

func (ru RegisterUser) ConfigValidation(v validate.Validation) {
	v.StopOnError = false
}
