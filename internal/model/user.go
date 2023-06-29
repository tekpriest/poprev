package model

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Base
	Username        string    `json:"username,omitempty"`
	Password        string    `json:"-"`
	Email           string    `json:"email,omitempty"`
	FirstName       string    `json:"first_name,omitempty"`
	LastName        string    `json:"last_name,omitempty"`
	AccountVerified bool      `json:"account_verified"`
	Tokens          []Token   `json:"tokens,omitempty"`
	Deposits        []Deposit `json:"deposits,omitempty"`
	Trades          []Trade   `json:"trades,omitempty"`
	Sales           []Sale    `json:"sales,omitempty"      gorm:"foreignKey:SellerID"`
}

var (
	ErrDuplicateEmail     = errors.New("email already exists")
	ErrDuplicateUsername  = errors.New("username already exists")
	ErrUnknownAccount     = errors.New("account does not exist")
	ErrMismatchedPassword = errors.New("incorrect email/password combination")
)

func (u *User) BeforeCreate(d *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	u.Email = strings.ToLower(u.Email)
	u.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().String())

	return
}

func (u *User) SetPassword(password string) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	u.Password = string(hash)

	return
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
