package response

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type HTTPResponse struct {
	Success bool        `json:"success"        default:"false"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"                 swaggertpe:"primitive,object"`
	Path    string      `json:"path,omitempty"`
} //	@Name	ErrorResponse

func BadRequestResponse(c *fiber.Ctx, message string, error interface{}) error {
	return c.Status(http.StatusBadRequest).JSON(&HTTPResponse{
		Success: false,
		Message: strings.ToTitle(message),
		Data:    error,
		Path:    c.Path(),
	})
}

func InternalServerErrorResponse(c *fiber.Ctx, message string, error interface{}) error {
	return c.Status(http.StatusInternalServerError).JSON(&HTTPResponse{
		Success: false,
		Message: strings.ToTitle(message),
		Data:    error,
		Path:    c.Path(),
	})
}

func NotFoundResponse(c *fiber.Ctx, message string, error interface{}) error {
	return c.Status(http.StatusNotFound).JSON(&HTTPResponse{
		Success: false,
		Message: strings.ToTitle(message),
		Data:    error,
		Path:    c.Path(),
	})
}

func UnauthorizedResponse(c *fiber.Ctx, message string, error interface{}) error {
	return c.Status(http.StatusUnauthorized).JSON(&HTTPResponse{
		Success: false,
		Message: strings.ToTitle(message),
		Data:    error,
		Path:    c.Path(),
	})
}

func CreatedResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(http.StatusCreated).JSON(&HTTPResponse{
		Success: true,
		Message: strings.ToTitle(message),
		Data:    data,
	})
}

func OkResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(http.StatusOK).JSON(&HTTPResponse{
		Success: true,
		Message: strings.ToTitle(message),
		Data:    data,
	})
}
