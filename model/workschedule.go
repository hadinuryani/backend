package model

import "time"

type WorkSchedule struct {
	ID         int       `json:"id"`
	EmployeeID int       `json:"employee_id"`
	StoreID    int       `json:"store_id"`
	ShiftID    int       `json:"shift_id"`
	Date       time.Time `json:"date"`
}
