package model

import "time"

type Attendance struct {
	ID             int        `json:"id"`
	EmployeeID     int        `json:"employee_id"`
	WorkScheduleID int        `json:"work_schedule_id"`
	CheckIn        time.Time  `json:"check_in"`
	CheckOut       *time.Time `json:"check_out"`
	Status         string     `json:"status" gorm:"type:enum('present','late','absent','permission','sick')"`
	Note           *string    `json:"note"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
