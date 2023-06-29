package user

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tekpriest/poprev/cmd/config"
	"github.com/tekpriest/poprev/cmd/database"
)

func Route(c *config.Config, r fiber.Router) {
	r = r.Group("/user")
	db := database.NewConnection(c)
	ctr := NewUserController(db)

	r.Get("/profile", ctr.Profile)
}
