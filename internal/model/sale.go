package model

type SaleStatus string

const (
	OnSale    SaleStatus = "on_sale"
	Completed SaleStatus = "completed"
	Cancelled SaleStatus = "cancelled"
)

type Sale struct {
	Base
	Quantity int        `json:"quantity,omitempty"`
	MinOrder int        `json:"min_order,omitempty"`
	MaxOrder int        `json:"max_order,omitempty"`
	Rate     float32    `json:"rate,omitempty"`
	Status   SaleStatus `json:"status,omitempty"`
	SellerID string     `json:"seller_id"`
	Trades   []Trade    `json:"trades,omitempty"`
}
