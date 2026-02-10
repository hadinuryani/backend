package model

type StoreStock struct {
	StoreID   int `json:"store_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
