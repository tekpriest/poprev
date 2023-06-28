package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/swagger"

	"github.com/tekpriest/poprev/cmd/config"
	"github.com/tekpriest/poprev/cmd/utils"
	_ "github.com/tekpriest/poprev/docs"
)

type Server interface {
	Serve() error
	fiber.Router
}

type server struct {
	c *config.Config
	*fiber.App
}

func (s *server) Serve() error {
	s.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status":    true,
			"message":   "Server up and running",
			"timestamp": time.Now().String(),
		})
	})

	utils.PanicOnError(
		s.Listen(fmt.Sprintf(":%d", s.c.Port)),
		"there was an issue running the server",
	)
	return nil
}

func NewRouter(c *config.Config) Server {
	r := fiber.New(fiber.Config{
		ErrorHandler: func(*fiber.Ctx, error) error {
			return nil
		},
	})

	// sensible defaults
	r.Use(requestid.New())
	r.Use(recover.New())
	r.Use(cors.New())
	r.Use(logger.New())
	// r.Use(limiter.New(limiter.Config{
	// 	Max:        10,
	// 	Expiration: 30 * time.Hour,
	// }))
	// r.Use(idempotency.New())
	r.Use(helmet.New())

	r.Get("/docs", swagger.HandlerDefault)

	return &server{c, r}
}
