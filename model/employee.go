package model

import "time"

type Employee struct {
	ID        int        `db:"id"`
	NIK       string     `db:"nik"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	Address   *string    `db:"address"`
	Phone     *string    `db:"phone"`
	HireDate  time.Time  `db:"hire_date"`
	EndHire   *time.Time `db:"end_hire"`
	Status    string     `db:"status"`
	CreatedAt time.Time  `db:"created_at"`
}