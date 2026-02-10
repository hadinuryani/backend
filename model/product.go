package model

type Product struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Category    *string `json:"category"`
	Unit        *string `json:"unit"`
}
