package model

type WarehouseStock struct {
	WarehouseID int `json:"warehouse_id"`
	ProductID   int `json:"product_id"`
	Quantity    int `json:"quantity"`
}
