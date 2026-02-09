package service

import (
	"backend-rems/config"
	"backend-rems/model"
	"backend-rems/model/response"
	"backend-rems/repository"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type HRDService struct {
	hrdRepo repository.HRDRepositoryInterface
}

func NewHRDService(hrdRepo repository.HRDRepositoryInterface) *HRDService {
	return &HRDService{hrdRepo: hrdRepo}
}

// business logic

func (s *HRDService) GetAllEmployees() ([]model.Employee, error) {
	return s.hrdRepo.GetAllEmployees()
}

func (s *HRDService) GetEmployeeByID(id int) (*response.EmployeeDetail, error) {
	rows, err := s.hrdRepo.GetEmployeeDetailRows(id)
	if err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return nil, errors.New("employee not found")
	}

	result := response.EmployeeDetail{
		ID:      rows[0].EmployeeID,
		NIK:     rows[0].NIK,
		Name:    rows[0].Name,
		Address: rows[0].Address,
		Status:  rows[0].Status,
		Stores:  []response.EmployeeStore{},
	}

	for _, r := range rows {
		if r.StoreID.Valid {
			result.Stores = append(result.Stores, response.EmployeeStore{
				StoreID:     int(r.StoreID.Int64),
				StoreName:   r.StoreName.String,
				RoleAtStore: r.RoleAtStore.String,
			})
		}
	}

	return &result, nil
}

func(s *HRDService)AddEmployees(employee model.Employee)(model.Employee,error){
	nik,err := s.generateNIK()
	if err !=nil{
		return model.Employee{},err
	}
	employee.NIK = nik
	return s.hrdRepo.AddEmployees(employee)
}

var hrdService *HRDService

func GetHRDService() *HRDService {
	if hrdService == nil {
		repo := repository.NewHRDRepository(config.DB)
		hrdService = NewHRDService(repo)
	}
	return hrdService
}

func(s *HRDService)UpdateEmployeeStatus(id int,status string)error {
	return s.hrdRepo.UpdateEmployeeStatus(id,status)
}

// membuat nik 
func (s *HRDService) generateNIK() (string, error) {
	today := time.Now().Format("20060102") // YYYYMMDD

	var lastNik string
	lastNik,err := s.hrdRepo.GetLastNikByDate(today)
	
	sequence := 1
	if err == nil {
		lastSeq, _ := strconv.Atoi(lastNik[8:])
		sequence = lastSeq + 1
	}

	return fmt.Sprintf("%s%04d", today, sequence), nil
}
