package handler

import (
	"backend-rems/helper"
	"backend-rems/model"
	"backend-rems/model/response"
	"backend-rems/service"
	"backend-rems/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /hrd/employees
func GetAllEmployees(ctx *gin.Context) {
	svc := service.GetHRDService()

	employees, err := svc.GetAllEmployees()
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch employees", err.Error())
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, "Employees fetched successfully", employees)
}

// GET /hrd/employees/:id
func GetEmployeeByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, "Invalid employee ID", nil)
		return
	}

	svc := service.GetHRDService()

	employee, err := svc.GetEmployeeByID(id)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusNotFound, err.Error(), nil)
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, "Employee detail fetched", employee)
}

// POST /api/hrd/employees
func AddEmployees(ctx *gin.Context) {
	var req validator.EmployeeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, "Validation Failed", err.Error())
		return
	}

	emp := model.Employee{
		Name:    req.Name,
		Address: &req.Address,
		Phone:   &req.Phone,
	}
	svc := service.GetHRDService()
	employee, err := svc.AddEmployees(emp)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to add Employe", err.Error())
		return
	}

	resp := response.EmployeeResponse{
		NIK:      employee.NIK,
		Name:     employee.Name,
		Address:  *employee.Address,
		Phone:    *employee.Phone,
		HireDate: employee.HireDate,
	}
	helper.SuccessResponse(ctx, http.StatusCreated, "Success to add employee", resp)
}

// PATCH /api/hrd/employees/:id/status
func UpdateEmployeeStatus(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id,err := strconv.Atoi(idParam)
	if err != nil {
		helper.ErrorResponse(ctx,http.StatusBadRequest,"Invalid employee ID",nil)
		return
	}

	var req validator.UpdateEmployeeStatusRequest
	if err := ctx.ShouldBindJSON(&req);err != nil{
		helper.ErrorResponse(ctx,http.StatusBadRequest,"Validation Failed",err.Error())
		return
	}

	svc := service.GetHRDService()
	if err := svc.UpdateEmployeeStatus(id, req.Status);err != nil{
		helper.ErrorResponse(ctx,http.StatusInternalServerError,"Failed to update status",err.Error())
		return
	}

	helper.SuccessResponse(ctx,http.StatusOK,"Employee status update",nil)
}