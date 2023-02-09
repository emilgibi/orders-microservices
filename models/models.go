package models

type Order struct {
	ID            int    `json:"id"`
	ProductName   string `json:"product_name"`
	OrderQuantity string `json:"order_quantity"`
}
