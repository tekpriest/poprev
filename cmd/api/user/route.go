package user

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tekpriest/poprev/cmd/config"
	"github.com/tekpriest/poprev/cmd/database"
	"github.com/tekpriest/poprev/internal/tokens"
)

func Route(c *config.Config, r fiber.Router) {
	r = r.Group("/user")
	db := database.NewConnection(c)
	ctr := NewUserController(db)

	authMiddleware := tokens.NewJwtMiddleware(c)
	r.Use(authMiddleware.JwtMiddleware)

	r.Get("/profile", ctr.Profile)
}
