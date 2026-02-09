package response

import "time"

type EmployeeDetail struct {
	ID      int             `json:"id"`
	NIK     string          `json:"nik"`
	Name    string          `json:"name"`
	Address string          `json:"address"`
	Status  string          `json:"status"`
	Stores  []EmployeeStore `json:"stores"`
}

type EmployeeStore struct {
	StoreID     int    `json:"store_id"`
	StoreName   string `json:"store_name"`
	RoleAtStore string `json:"role_at_store"`
}

type EmployeeResponse struct {
	NIK      string    `json:"nik"`
	Name     string    `json:"name"`
	Address  string    `json:"address,omitempty"`
	Phone    string    `json:"phone,omitempty"`
	HireDate time.Time `json:"hire_date"`
}
