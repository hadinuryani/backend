package model

type StoreEmployee struct {
	StoreID    int    `db:"store_id"`
	EmployeeID int    `db:"employee_id"`
	RoleAtStore string `db:"role_at_store"`
}
