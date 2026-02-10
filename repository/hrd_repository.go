package repository

import (
	"backend-rems/model"
	"database/sql"

)

type HRDRepositoryInterface interface {
	GetAllEmployees() ([]model.Employee, error)
	GetEmployeeDetailRows(id int) ([]EmployeeDetailRow, error)
	AddEmployees(employee model.Employee) (model.Employee, error)
	GetLastNikByDate(date string) (string, error)
	UpdateEmployeeStatus(id int, status string) error
	SetRole(role model.Role) error
	GetRoles() ([]model.Role,error)
	UpdateRole(id int,data model.Role) error
	DeleteRole(id int) error
	CreateEmployee(emp model.Employee) (model.Employee, error)
	AssignEmployeeToStore(storeID int,employeeID int,roleAtStore string,) error 
	GetAllStores() ([]model.Store, error)
}

type HRDRepository struct {
	db *sql.DB
}

func NewHRDRepository(db *sql.DB) *HRDRepository {
	return &HRDRepository{db: db}
}

func (r *HRDRepository) SetRole(role model.Role) error {
	query := `INSERT INTO roles (role_name,base_salary)
	VALUES (?, ?)`

	_, err := r.db.Exec(query,role.RoleName,role.BaseSalary)
	if err != nil {
		return err
	}

	return nil
}

func (r *HRDRepository) UpdateRole(id int, data model.Role) error {
	query := `
		UPDATE roles 
		SET role_name = ?, base_salary = ? 
		WHERE id = ?
	`

	result, err := r.db.Exec(query, data.RoleName, data.BaseSalary, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}


func (r *HRDRepository) DeleteRole(id int) error {
	query := `DELETE FROM roles WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}


func (r *HRDRepository) GetRoles() ([]model.Role, error) {
	query := `SELECT id, role_name, base_salary FROM roles`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []model.Role
	for rows.Next() {
		var role model.Role
		if err := rows.Scan(
			&role.ID,
			&role.RoleName,
			&role.BaseSalary,
		); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}


func (r *HRDRepository) GetAllEmployees() ([]model.Employee, error) {
	query := `
		SELECT id, nik, name, phone, hire_date, status
		FROM employees
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []model.Employee

	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(
			&e.ID,
			&e.NIK,
			&e.Name,
			&e.Phone,
			&e.HireDate,
			&e.Status,
		); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}

	return employees, nil
}

func (r *HRDRepository) GetEmployeeDetailRows(id int) ([]EmployeeDetailRow, error) {
	query := `
		SELECT 
			e.id,
			e.nik,
			e.name,
			e.address,
			e.status,
			st.id,
			st.store_name,
			se.role_at_store
		FROM employees e
		LEFT JOIN stores_employees se ON e.id = se.employee_id
		LEFT JOIN stores st ON st.id = se.store_id
		WHERE e.id = ?
	`

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []EmployeeDetailRow

	for rows.Next() {
		var row EmployeeDetailRow
		if err := rows.Scan(
			&row.EmployeeID,
			&row.NIK,
			&row.Name,
			&row.Address,
			&row.Status,
			&row.StoreID,
			&row.StoreName,
			&row.RoleAtStore,
		); err != nil {
			return nil, err
		}
		results = append(results, row)
	}

	return results, nil
}

func (r *HRDRepository) AddEmployees(employee model.Employee) (model.Employee, error) {
	query := `
			INSERT INTO employees
			(nik,name,address,phone,hire_date)
			VALUES(?,?,?,?,NOW())`

	result, err := r.db.Exec(
		query,
		employee.NIK,
		employee.Name,
		employee.Address,
		employee.Phone,
	)
	if err != nil {
		return model.Employee{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return model.Employee{}, err
	}
	err = r.db.QueryRow(
		"SELECT nik,name,phone,hire_date FROM employees WHERE id = ?",
		id).Scan(
		&employee.NIK,
		&employee.Name,
		&employee.Phone,
		&employee.HireDate,
	)
	if err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}

func (r *HRDRepository) GetLastNikByDate(date string) (string, error) {
	var nik string

	query := `SELECT nik
			FROM employees
			WHERE nik LIKE ?
			ORDER BY nik DESC
			LIMIT 1	`

	err := r.db.QueryRow(query, date+"%").Scan(&nik)
	if err != nil {
		return "", err
	}
	return nik, nil
}

func (r *HRDRepository) UpdateEmployeeStatus(id int, status string) error {
	query := `UPDATE employees SET status = ? WHERE id = ?`

	_, err := r.db.Exec(query, status, id)
	return err
}

func (r *HRDRepository) CreateEmployee(emp model.Employee) (model.Employee, error) {
	query := `
		INSERT INTO employees (nik, name, email, address, phone, hire_date, status)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.Exec(
		query,
		emp.NIK,
		emp.Name,
		emp.Email,
		emp.Address,
		emp.Phone,
		emp.HireDate,
		emp.Status,
	)
	if err != nil {
		return emp, err
	}

	id, _ := result.LastInsertId()
	emp.ID = int(id)
	return emp, nil
}

func (r *HRDRepository) AssignEmployeeToStore(
	storeID int,
	employeeID int,
	roleAtStore string,
) error {

	query := `
		INSERT INTO stores_employees (store_id, employee_id, role_at_store)
		VALUES (?, ?, ?)
	`
	_, err := r.db.Exec(query, storeID, employeeID, roleAtStore)
	return err
}

func (r *HRDRepository) GetAllStores() ([]model.Store, error) {
	query := `SELECT id, store_name, address FROM stores ORDER BY store_name`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stores []model.Store
	for rows.Next() {
		var s model.Store
		if err := rows.Scan(&s.ID, &s.StoreName, &s.Address); err != nil {
			return nil, err
		}
		stores = append(stores, s)
	}

	return stores, nil
}
