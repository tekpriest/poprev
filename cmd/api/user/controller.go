package user

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tekpriest/poprev/cmd/database"
	response "github.com/tekpriest/poprev/internal/types"
)

type UserController interface {
	Profile(ctx *fiber.Ctx) error
}

type controller struct {
	s UserService
}

// Profile implements UserController.
func (c *controller) Profile(ctx *fiber.Ctx) error {
	userID := ctx.Get("user_id")
	if userID == "" {
		return response.UnauthorizedResponse(ctx, "invalid user session. please login again", nil)
	}

	user, err := c.s.FindById(userID)
	if err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err)
	}

	return response.OkResponse(ctx, "profile fetched successfully", user)
}

func NewUserController(db database.DatabaseConnection) UserController {
	s := NewUserService(db)
	return &controller{s: s}
}
