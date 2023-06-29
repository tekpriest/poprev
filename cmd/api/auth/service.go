package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/tekpriest/poprev/cmd/api/user"
	"github.com/tekpriest/poprev/cmd/config"
	"github.com/tekpriest/poprev/cmd/database"
	"github.com/tekpriest/poprev/internal/model"
	"github.com/tekpriest/poprev/internal/tokens"
)

var ctx = context.TODO()

type AuthService interface {
	Register(data user.RegisterUser) (AuthData, error)
	VerifyAccount(code string) (AuthData, error)
	Login(email, password string) (AuthData, error)
	CreatePasswordRequest(email string) (string, error)
	ResetPassword(data ResetPasswordData) (AuthData, error)
}

type service struct {
	db  *gorm.DB
	rdb *redis.Client
	c   *config.Config
	us  user.UserService
}

// CreatePasswordRequest implements AuthService.
func (s *service) CreatePasswordRequest(email string) (string, error) {
	user, err := s.us.FindByEmail(email)
	if err != nil {
		return "", errors.New(model.ErrUnknownAccount.Error())
	}

	otp, err := tokens.GenerateOTP()
	if err != nil {
		return "", err
	}

	if err := s.rdb.
		Set(
			ctx,
			fmt.Sprintf("reset_%s", otp.Hash),
			user.ID,
			time.Duration(time.Hour*24)).
		Err(); err != nil {
		return "", err
	}

	return otp.Secret, nil
}

// ResetPassword implements AuthService.
func (s *service) ResetPassword(data ResetPasswordData) (AuthData, error) {
	var auth AuthData

	userID := s.rdb.Get(ctx, fmt.Sprintf("reset_%s", data.Code))
	if userID == nil {
		return AuthData{}, errors.New("invalid reset password code. please try again")
	}

	user, err := s.us.FindById(userID.String())
	if err != nil {
		return AuthData{}, errors.New(model.ErrUnknownAccount.Error())
	}

	user, err = s.us.UpdatePassword(user.ID, data.Password)
	if err != nil {
		return AuthData{}, errors.New(model.ErrUnknownAccount.Error())
	}

	auth.User = *user
	auth.Token, err = s.createAuthToken(user.ID)
	if err != nil {
		return AuthData{}, err
	}

	return auth, nil
}

// Login implements AuthService.
func (s *service) Login(email, password string) (AuthData, error) {
	var auth AuthData

	user, err := s.us.FindByEmail(email)
	if err != nil {
		return AuthData{}, errors.New(model.ErrUnknownAccount.Error())
	}

	if valid := user.ComparePassword(password); !valid {
		return AuthData{}, errors.New(model.ErrMismatchedPassword.Error())
	}

	auth.User = *user
	auth.Token, err = s.createAuthToken(user.ID)
	if err != nil {
		return AuthData{}, err
	}

	return auth, nil
}

// Register implements AuthService.
func (s *service) Register(data user.RegisterUser) (AuthData, error) {
	var auth AuthData
	user, _ := s.us.FindByEmail(strings.ToLower(data.Email))
	if user != nil {
		return AuthData{}, errors.New(model.ErrDuplicateEmail.Error())
	}

	user, _ = s.us.FindByUsername(data.Username)
	if user != nil {
		return AuthData{}, errors.New(model.ErrDuplicateUsername.Error())
	}

	user, err := s.us.CreateUser(data)
	if err != nil {
		return AuthData{}, err
	}

	if err != nil {
		return AuthData{}, err
	}

	otp, err := tokens.GenerateOTP()
	if err != nil {
		return AuthData{}, err
	}

	if err := s.rdb.
		Set(
			ctx,
			fmt.Sprintf("verify_%s", otp.Hash),
			user.ID,
			time.Duration(time.Hour*24)).
		Err(); err != nil {
		return AuthData{}, err
	}

	// TODO: to send activation email to user
	// but for purpose of this test, I am returning it
	auth.ActivationCode = otp.Secret

	return auth, nil
}

// VerifyAccount implements AuthService.
func (s *service) VerifyAccount(code string) (AuthData, error) {
	var auth AuthData
	userID := s.rdb.Get(ctx, fmt.Sprintf("verify_%s", code))
	if userID == nil {
		return AuthData{}, errors.New("invalid verification code please try again")
	}

	user, err := s.us.VerifyAccount(userID.String())
	if err != nil {
		return AuthData{}, err
	}

	auth.Token, _ = s.createAuthToken(user.ID)
	auth.User = *user

	return auth, nil
}

func (s *service) createAuthToken(userID string) (string, error) {
	secret := []byte(s.c.JwtSecret)
	claims := AuthTokenData{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	tokenData := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenData.SignedString(secret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func NewAuthService(c *config.Config, db database.DatabaseConnection) AuthService {
	us := user.NewUserService(db)

	return &service{
		db:  db.GetDB(),
		rdb: db.GetRBD(),
		c:   c,
		us:  us,
	}
}
