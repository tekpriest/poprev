package project

import (
	"github.com/gofiber/fiber/v2"

	"github.com/tekpriest/poprev/cmd/config"
	"github.com/tekpriest/poprev/cmd/database"
)

func Route(c *config.Config, r fiber.Router) {
	r = r.Group("/projects")
	db := database.NewConnection(c)
	ctr := NewProjectController(db)

	r.Post("/create", ctr.CreateProject)
	r.Get("/", ctr.FetchAllProjects)
	r.Get("/:id", ctr.FetchAProject)
}
