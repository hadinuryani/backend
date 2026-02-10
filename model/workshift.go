package model

import "time"

type WorkShift struct {
	ID        int       `json:"id"`
	ShiftName string    `json:"shift_name"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
