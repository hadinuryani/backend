package validator


type EmployeeRequest struct {
	Name       string `json:"name" binding:"required,min=3,max=100"`
	Email      string `json:"email" binding:"required,email"`
	Address    string `json:"address" binding:"required"`
	Phone      string `json:"phone" binding:"required,min=10,max=15"`
	StartDate  string `json:"start_kontrak" binding:"required,datetime=2006-01-02"`
	StoreID    int    `json:"store_id" binding:"required"`
}



type RoleRequest struct {
	ID         int    `json:"id"`
	RoleName   string `json:"role_name" binding:"required"`
	BaseSalary int64  `json:"base_salary" binding:"required" `
}

// type EmployeeRequest struct {
// 	Name     string     `json:"name"`
// 	Email    string     `json:"email"`
// 	Address  *string    `json:"address"`
// 	Phone    *string    `json:"phone"`
// 	Contract string 	`json:"contract"`
// 	HireDate time.Time  `json:"hire_date"`
// 	EndHire  *time.Time `json:"end_contract"`
// 	Status   string     `json:"status"`
// 	Gaji     int        ``
// }
