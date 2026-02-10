package model

type Role struct {
	ID         int     `json:"id"`
	RoleName   string  `json:"role_name"`
	BaseSalary int64 `json:"base_salary"`
}
