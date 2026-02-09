package repository

import "database/sql"

// untuk query hasil join
type EmployeeDetailRow struct {
	EmployeeID  int
	NIK         string
	Name        string
	Address     string
	Status      string
	StoreID     sql.NullInt64
	StoreName   sql.NullString
	RoleAtStore sql.NullString
}

