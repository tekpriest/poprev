package token

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tekpriest/poprev/cmd/config"
	"github.com/tekpriest/poprev/cmd/database"
	"github.com/tekpriest/poprev/internal/tokens"
)

func Route(c *config.Config, r fiber.Router) {
	r = r.Group("/tokens")
	db := database.NewConnection(c)
	ctr := NewTokenController(db)

	authMiddleware := tokens.NewJwtMiddleware(c)
	r.Use(authMiddleware.JwtMiddleware)

	r.Get("/", ctr.FetchAllTokens)
	r.Post("/buy", ctr.BuyToken)
	r.Post("/sell", ctr.CreateSale)
	r.Get("/sales", ctr.FetchAllSales)
	r.Get("/sales/:id", ctr.FetchASale)
}
