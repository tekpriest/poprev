package auth

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tekpriest/poprev/cmd/config"
	"github.com/tekpriest/poprev/cmd/database"
)

func Route(c *config.Config, r fiber.Router) {
	r = r.Group("/auth")
	db := database.NewConnection(c)
	ctr := NewAuthController(c, db)

	// r.Use(internal.ValidateBody)

	r.Post("/register", ctr.Register)
	r.Post("/login", ctr.Login)
	r.Post("/password/request", ctr.RequestPassword)
	r.Post("/password/reset", ctr.ResetPassword)
}
