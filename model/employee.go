package model

import "time"

type Employee struct {
	ID       int       `json:"id"`
	NIK      string    `json:"nik"`
	Name     string    `json:"name"`
	Address  *string   `json:"address"`
	Phone    *string   `json:"phone"`
	HireDate time.Time `json:"hire_date"`
	Status   string    `json:"status"`
}
