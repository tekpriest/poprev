package internal

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	response "github.com/tekpriest/poprev/internal/types"
)

func ValidateBody(c *fiber.Ctx) error {
	validater := validator.New()
	var body interface{}
	var errors []string

	if err := c.BodyParser(&body); err != nil {
		return response.BadRequestResponse(c, err.Error(), err)
	}

	if err := validater.Struct(body); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			msg := fmt.Sprintf("%s is a %s field", err.Field(), err.Tag())
			errors = append(errors, msg)
		}
	}

	if len(errors) > 0 {
		return response.BadRequestResponse(c, "there are missing required fields", errors)
	}

	return c.Next()
}
