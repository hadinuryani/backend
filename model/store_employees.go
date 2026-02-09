package model

type StoreEmployees struct {
	StoreID     int    `json:"store_id"`
	EmployeeId  int    `json:"employe_id"`
	RoleAtStore string `json:"role_at_store"`
}
