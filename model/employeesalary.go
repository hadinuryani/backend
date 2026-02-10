package model

import "time"

type EmployeeSalary struct {
	ID          int       `json:"id"`
	EmployeeID  int       `json:"employee_id"`
	ComponentID int       `json:"component_id"`
	Period      time.Time `json:"period"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
}
