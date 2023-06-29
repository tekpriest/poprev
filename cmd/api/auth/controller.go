package auth

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"

	"github.com/tekpriest/poprev/cmd/api/user"
	"github.com/tekpriest/poprev/cmd/config"
	"github.com/tekpriest/poprev/cmd/database"
	response "github.com/tekpriest/poprev/internal/types"
)

type AuthController interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	RequestPassword(ctx *fiber.Ctx) error
	ResetPassword(ctx *fiber.Ctx) error
	VerifyAccount(ctx *fiber.Ctx) error
}

type controller struct {
	s AuthService
}

// VerifyAccount implements AuthController.
func (c *controller) VerifyAccount(ctx *fiber.Ctx) error {
	var data VerifyAccountData

	json.Unmarshal(ctx.Body(), &data)

	authData, err := c.s.VerifyAccount(data.Code)
	if err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err.Error())
	}

	return response.OkResponse(ctx, "account verified successfully", authData)
}

// Login implements AuthController.
func (c *controller) Login(ctx *fiber.Ctx) error {
	var data LoginRequestData

	json.Unmarshal(ctx.Body(), &data)

	authData, err := c.s.Login(data.Email, data.Password)
	if err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err.Error())
	}

	return response.OkResponse(ctx, "logged in successfully", authData)
}

// Register implements AuthController.
func (c *controller) Register(ctx *fiber.Ctx) error {
	var data user.RegisterUser

	json.Unmarshal(ctx.Body(), &data)

	authData, err := c.s.Register(data)
	if err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err.Error())
	}

	return response.CreatedResponse(ctx, "account created successfully", authData)
}

// RequestPassword implements AuthController.
func (c *controller) RequestPassword(ctx *fiber.Ctx) error {
	var data PasswordRequestData

	json.Unmarshal(ctx.Body(), &data)

	code, err := c.s.CreatePasswordRequest(data.Email)
	if err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err.Error())
	}

	return response.OkResponse(ctx, "password reset token created", map[string]string{
		"code": code,
	})
}

// ResetPassword implements AuthController.
func (c *controller) ResetPassword(ctx *fiber.Ctx) error {
	var data ResetPasswordData
	json.Unmarshal(ctx.Body(), &data)

	authData, err := c.s.ResetPassword(data)
	if err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err.Error())
	}

	return response.OkResponse(ctx, "password reset successfully", authData)
}

func NewAuthController(c *config.Config, db database.DatabaseConnection) AuthController {
	s := NewAuthService(c, db)

	return &controller{s: s}
}
