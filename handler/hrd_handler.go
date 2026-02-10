package handler

import (
	"backend-rems/helper"
	"backend-rems/model"
	"backend-rems/service"
	"backend-rems/validator"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetRole(ctx *gin.Context) {
	var req validator.RoleRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, "Error cik salah validasine", err.Error())
		return
	}

	roles := model.Role{
		RoleName:   req.RoleName,
		BaseSalary: req.BaseSalary,
	}
	svc := service.GetHRDService()
	if err := svc.SetRole(roles); err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed insert roles", err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, "Role berhasil di tambahkan ", &roles)
}

func UpdateRole(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid role ID", nil)
		return
	}

	var req validator.RoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, "Validation error", err.Error())
		return
	}
	fmt.Println(&req)

	role := model.Role{
		RoleName:   req.RoleName,
		BaseSalary: req.BaseSalary,
	}

	svc := service.GetHRDService()
	if err := svc.UpdateRole(id, role); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			helper.ErrorResponse(ctx, http.StatusNotFound, "Role not found", nil)
			return
		}

		helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed update role", err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, "Role berhasil diupdate", nil)
}


func GetRole(ctx *gin.Context) {
	svc := service.GetHRDService()

	roles, err := svc.GetRoles()
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed get roles", err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, "Fetch roles successfully", roles)
}



func DeleteRole(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid role ID", nil)
		return
	}

	svc := service.GetHRDService()
	if err := svc.DeleteRole(id); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			helper.ErrorResponse(ctx, http.StatusNotFound, "Role not found", nil)
			return
		}

		helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed delete role", err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, "Role berhasil dihapus", nil)
}

func AddEmployees(ctx *gin.Context) {
	var req validator.EmployeeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, "Validation Error", err.Error())
		return
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid Date", "Format YYYY-MM-DD")
		return
	}

	employee := model.Employee{
		Name:     req.Name,
		Email:    req.Email,
		Address:  &req.Address,
		Phone:    &req.Phone,
		HireDate: startDate,
		Status:   "active",
	}

	svc := service.GetHRDService()
	emp, err := svc.AddEmployee(employee, req.StoreID,req.Role)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed add employee", err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusCreated, "Employee created", emp)
}

func GetStores(ctx *gin.Context) {
	svc := service.GetHRDService()

	stores, err := svc.GetStores()
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed get stores", err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, "Fetch stores success", stores)
}






// // GET /hrd/employees
// func GetAllEmployees(ctx *gin.Context) {
// 	svc := service.GetHRDService()

// 	employees, err := svc.GetAllEmployees()
// 	if err != nil {
// 		helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch employees", err.Error())
// 		return
// 	}

// 	helper.SuccessResponse(ctx, http.StatusOK, "Employees fetched successfully", employees)
// }

// // GET /hrd/employees/:id
// func GetEmployeeByID(ctx *gin.Context) {
// 	idParam := ctx.Param("id")

// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid employee ID", nil)
// 		return
// 	}

// 	svc := service.GetHRDService()

// 	employee, err := svc.GetEmployeeByID(id)
// 	if err != nil {
// 		helper.ErrorResponse(ctx, http.StatusNotFound, err.Error(), nil)
// 		return
// 	}

// 	helper.SuccessResponse(ctx, http.StatusOK, "Employee detail fetched", employee)
// }

// // POST /api/hrd/employees
// func AddEmployees(ctx *gin.Context) {
// 	var req validator.EmployeeRequest

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		helper.ErrorResponse(ctx, http.StatusBadRequest, "Validation Failed", err.Error())
// 		return
// 	}

// 	emp := model.Employee{
// 		Name:    req.Name,
// 		NIK: req.NIK,
// 		Email : req.Email,
// 		HireDate: req.HireDate,
// 		TokoId : req.TokoId,
// 		Address: &req.Address,
// 		Phone:   &req.Phone,
// 	}
// 	svc := service.GetHRDService()
// 	employee, err := svc.AddEmployees(emp)
// 	if err != nil {
// 		helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to add Employe", err.Error())
// 		return
// 	}

// 	resp := response.EmployeeResponse{
// 		Name:     employee.Name,
// 		NIK:      employee.NIK,
// 		Address:  *employee.Address,
// 		Phone:    *employee.Phone,
// 		HireDate: employee.HireDate,
// 	}
// 	helper.SuccessResponse(ctx, http.StatusCreated, "Success to add employee", resp)
// }

// // PATCH /api/hrd/employees/:id/status
// func UpdateEmployeeStatus(ctx *gin.Context) {
// 	idParam := ctx.Param("id")
// 	id,err := strconv.Atoi(idParam)
// 	if err != nil {
// 		helper.ErrorResponse(ctx,http.StatusBadRequest,"Invalid employee ID",nil)
// 		return
// 	}

// 	var req validator.UpdateEmployeeStatusRequest
// 	if err := ctx.ShouldBindJSON(&req);err != nil{
// 		helper.ErrorResponse(ctx,http.StatusBadRequest,"Validation Failed",err.Error())
// 		return
// 	}

// 	svc := service.GetHRDService()
// 	if err := svc.UpdateEmployeeStatus(id, req.Status);err != nil{
// 		helper.ErrorResponse(ctx,http.StatusInternalServerError,"Failed to update status",err.Error())
// 		return
// 	}

// 	helper.SuccessResponse(ctx,http.StatusOK,"Employee status update",nil)
// }
