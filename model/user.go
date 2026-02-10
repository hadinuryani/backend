package model

type User struct {
	ID           int    `json:"id"`
	EmployeeID   int    `json:"employee_id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	IsActive     bool   `json:"is_active"`
}
