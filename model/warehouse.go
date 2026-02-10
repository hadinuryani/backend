package model

type Warehouse struct {
	ID            int     `json:"id"`
	WarehouseName string  `json:"warehouse_name"`
	Address       *string `json:"address"`
}
