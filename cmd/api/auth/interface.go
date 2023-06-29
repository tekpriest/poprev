package auth

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/tekpriest/poprev/internal/model"
)

type AuthData struct {
	User           model.User `json:"user,omitempty"`
	Token          string     `json:"token,omitempty"`
	ActivationCode string     `json:"code,omitempty"`
}

type AuthTokenData struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

type LoginRequestData struct {
	Email    string `json:"email"    validate:"required|email"`
	Password string `json:"password" validate:"required"`
}

type PasswordRequestData struct {
	Email string `json:"email" validate:"required|email"`
}

type ResetPasswordData struct {
	Password string `json:"password" validate:"required"`
	Code     string `json:"code"     validate:"required"`
}
