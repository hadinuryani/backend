package model

type Store struct {
	ID        int     `json:"id"`
	StoreName string  `json:"store_name"`
	Address   *string `json:"address"`
}
