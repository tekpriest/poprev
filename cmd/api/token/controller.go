package token

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"

	"github.com/tekpriest/poprev/cmd/database"
	response "github.com/tekpriest/poprev/internal/types"
)

type TokenController interface {
	CreateSale(ctx *fiber.Ctx) error
	FetchASale(ctx *fiber.Ctx) error
	FetchAllSales(ctx *fiber.Ctx) error
	BuyToken(ctx *fiber.Ctx) error
	FetchAllTokens(ctx *fiber.Ctx) error
}

type controller struct {
	s TokenService
}

// BuyToken implements TokenController.
func (c *controller) BuyToken(ctx *fiber.Ctx) error {
	var data BuyTokenData

	if err := ctx.BodyParser(&data); err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err)
	}

	v := validate.New(data)

	if ok := v.Validate(); !ok {
		return response.BadRequestResponse(ctx, "there are missing required fields", v.Errors.All())
	}

	if err := v.BindSafeData(&data); err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err)
	}

	userID := ctx.Get("user_id")
	if userID == "" {
		return response.UnauthorizedResponse(ctx, "invalid user session. please login again", nil)
	}

	result, err := c.s.BuyToken(userID, data)
	if err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err.Error())
	}

	return response.CreatedResponse(ctx, "token purchansed successfully", result)
}

// FetchAllSales implements TokenController.
func (c *controller) FetchAllSales(ctx *fiber.Ctx) error {
	var query QueryUserSalesData

	userID := ctx.Get("user_id")
	if userID == "" {
		return response.UnauthorizedResponse(ctx, "invalid user session. please login again", nil)
	}

	result, err := c.s.FetchSales(userID, query)
	if err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err.Error())
	}

	return response.OkResponse(ctx, "sales fetched successfully", result)
}

// FetchAllTokens implements TokenController.
func (c *controller) FetchAllTokens(ctx *fiber.Ctx) error {
	var query QueryUserTokensData

	userID := ctx.Get("user_id")
	if userID == "" {
		return response.UnauthorizedResponse(ctx, "invalid user session. please login again", nil)
	}

	result, err := c.s.FetchAllTokens(userID, query)
	if err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err.Error())
	}

	return response.OkResponse(ctx, "tokens fetched successfully", result)
}

// CreateSale implements TokenController.
func (c *controller) CreateSale(ctx *fiber.Ctx) error {
	var data CreateSaleData

	if err := ctx.BodyParser(&data); err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err)
	}

	v := validate.New(data)

	if ok := v.Validate(); !ok {
		return response.BadRequestResponse(ctx, "there are missing required fields", v.Errors.All())
	}

	if err := v.BindSafeData(&data); err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err)
	}

	userID := ctx.Get("user_id")
	if userID == "" {
		return response.UnauthorizedResponse(ctx, "invalid user session. please login again", nil)
	}

	sale, err := c.s.CreateSale(userID, data)
	if err != nil {
		return response.BadRequestResponse(ctx, err.Error(), err.Error())
	}

	return response.CreatedResponse(ctx, "sale created successfully", sale)
}

// FetchASale implements TokenController.
func (c *controller) FetchASale(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")

	if ID == "" {
		return response.BadRequestResponse(ctx, "empty params in path", nil)
	}

	sale, err := c.s.FetchSale(ID)
	if err != nil {
		return response.NotFoundResponse(ctx, err.Error(), err)
	}

	return response.OkResponse(ctx, "fetched sale details", sale)
}

func NewTokenController(db database.DatabaseConnection) TokenController {
	s := NewTokenService(db)
	return &controller{s: s}
}
