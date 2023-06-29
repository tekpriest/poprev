package project

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"

	"github.com/tekpriest/poprev/cmd/database"
	response "github.com/tekpriest/poprev/internal/types"
)

type ProjectController interface {
	CreateProject(ctx *fiber.Ctx) error
	FetchAllProjects(ctx *fiber.Ctx) error
	FetchAProject(ctx *fiber.Ctx) error
	FetchAllTransactions(ctx *fiber.Ctx) error
}

type controller struct {
	s ProjectService
}

// FetchAllTransactions implements ProjectController.
func (c *controller) FetchAllTransactions(ctx *fiber.Ctx) error {
	var query QueryTransactionsData

	projectID := ctx.Params("id")
	if projectID == "" {
		return response.BadRequestResponse(ctx, "empty id in params path", nil)
	}

	if err := ctx.QueryParser(&query); err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err.Error())
	}

	result, err := c.s.FetchAllProjectTransactions(projectID, query)
	if err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err.Error())
	}

	return response.OkResponse(ctx, "fetched all transactions successfully", result)
}

// CreateProject implements ProjectController.
func (c *controller) CreateProject(ctx *fiber.Ctx) error {
	var data CreateProjectData

	json.Unmarshal(ctx.Body(), &data)

	project, err := c.s.CreateProject(data)
	if err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err.Error())
	}

	return response.CreatedResponse(ctx, "project created successfully", project)
}

// FetchAProject implements ProjectController.
func (c *controller) FetchAProject(ctx *fiber.Ctx) error {
	projectID := ctx.Params("id")
	if projectID == "" {
		return response.BadRequestResponse(ctx, "empty id in params path", nil)
	}

	project, err := c.s.FetchProject(projectID)
	if err != nil {
		return response.NotFoundResponse(ctx, err.Error(), err)
	}

	return response.OkResponse(ctx, "project details found successfully", project)
}

// FetchAllProjects implements ProjectController.
func (c *controller) FetchAllProjects(ctx *fiber.Ctx) error {
	var query QueryProjectData

	if err := ctx.QueryParser(&query); err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err.Error())
	}

	result, err := c.s.ListProjects(query)
	if err != nil {
		return response.NotFoundResponse(ctx, err.Error(), err)
	}

	return response.OkResponse(ctx, "projects fetched successfully", result)
}

func NewProjectController(db database.DatabaseConnection) ProjectController {
	s := NewProjectService(db)

	return &controller{s: s}
}
