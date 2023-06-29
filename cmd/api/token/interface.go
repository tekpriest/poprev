package token

import (
	"github.com/tekpriest/poprev/internal/model"
	"github.com/tekpriest/poprev/internal/query"
)

type BuyTokenData struct {
	ProjectID string `json:"project_id,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
}

type QueryUserTokensData struct {
	Limit int64 `query:"limit"`
	Page  int64 `query:"page"`
}

type QueryUserSalesData struct {
	Limit int64 `query:"limit"`
	Page  int64 `query:"page"`
}

type FetchTokensData struct {
	Tokens []model.Token    `json:"tokens"`
	Meta   query.Pagination `json:"meta,omitempty"`
}

type CreateSaleData struct {
	TokenID  string `json:"token_id,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
	MinOrder int    `json:"min_order,omitempty"`
	MaxOrder int    `json:"max_order,omitempty"`
}

type FetchSalesData struct {
	Sales []model.Sale     `json:"tokens"`
	Meta  query.Pagination `json:"meta,omitempty"`
}
