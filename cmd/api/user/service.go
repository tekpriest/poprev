package user

import (
	"errors"
	"strings"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/tekpriest/poprev/cmd/database"
	"github.com/tekpriest/poprev/internal/model"
)

type UserService interface {
	CreateUser(data RegisterUser) (*model.User, error)
	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	Update(id string, data UpdateUser) (*model.User, error)
	VerifyAccount(id string) (*model.User, error)
	UpdatePassword(id, password string) (*model.User, error)
}

type service struct {
	db  *gorm.DB
	rdb *redis.Client
}

// UpdatePassword implements UserService.
func (s *service) UpdatePassword(id string, password string) (*model.User, error) {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	if err := s.db.
		Table("users").
		Where("id = ?", id).
		Update("password", newPassword).Error; err != nil {
		return nil, err
	}

	return s.FindById(id)
}

// FindByUsername implements UserService.
func (s *service) FindByUsername(username string) (*model.User, error) {
	var user model.User

	if err := s.db.Table("users").
		First(&user, "username = ?", username).
		Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser implements UserService.
func (s *service) CreateUser(data RegisterUser) (*model.User, error) {
	user := &model.User{
		Username:  data.Username,
		Email:     data.Email,
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}

	if err := user.SetPassword(data.Password); err != nil {
		return nil, err
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return s.FindById(user.ID)
}

// FindByEmail implements UserService.
func (s *service) FindByEmail(email string) (*model.User, error) {
	var user model.User

	if err := s.db.Table("users").
		First(&user, "email = ?", strings.ToLower(email)).
		Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// FindById implements UserService.
func (s *service) FindById(id string) (*model.User, error) {
	var user model.User

	if err := s.db.Table("users").
		First(&user, "id = ?", id).
		Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Update implements UserService.
func (s *service) Update(id string, data UpdateUser) (*model.User, error) {
	_, err := s.FindById(id)
	if err != nil {
		return nil, err
	}

	if err := s.db.
		Table("users").
		Where("id = ?", id).
		Updates(data).
		Error; err != nil {
		return nil, err
	}

	return s.FindById(id)
}

func (s *service) VerifyAccount(ID string) (*model.User, error) {
	user, err := s.FindById(ID)
	if err != nil {
		return nil, err
	}

	if user.AccountVerified {
		return nil, errors.New("account has already been verified")
	}

	if err := s.db.
		Table("users").
		Where("id = ?", ID).
		Update("account_verified", true).Error; err != nil {
		return nil, err
	}

	return s.FindById(ID)
}

func NewUserService(db database.DatabaseConnection) UserService {
	return &service{
		db:  db.GetDB(),
		rdb: db.GetRBD(),
	}
}
