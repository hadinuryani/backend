package service

import (
	"backend-rems/config"
	"backend-rems/model"
	"backend-rems/repository"
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

func(s *HRDService)SetRole(role model.Role)error{
	return s.hrdRepo.SetRole(role)
}

func(s *HRDService)UpdateRole(id int,role model.Role)error{
	return s.hrdRepo.UpdateRole(id,role)
}
func(s *HRDService)DeleteRole(id int)error{
	return s.hrdRepo.DeleteRole(id)
}

func(s *HRDService)GetRoles()([]model.Role,error){
	return s.hrdRepo.GetRoles()
}

func (s *HRDService) GetAllEmployees() ([]model.Employee, error) {
	return s.hrdRepo.GetAllEmployees()
}
func (s *HRDService) GetStores() ([]model.Store, error) {
	return s.hrdRepo.GetAllStores()
}






func (s *HRDService) AddEmployee(employee model.Employee,storeID int,) (model.Employee, error) {
	nik, err := s.generateNIK()
	if err != nil {
		return model.Employee{}, err
	}
	employee.NIK = nik

	// 1. insert employee
	emp, err := s.hrdRepo.CreateEmployee(employee)
	if err != nil {
		return model.Employee{}, err
	}

	// 2. assign ke store
	err = s.hrdRepo.AssignEmployeeToStore(
		storeID,
		emp.ID,
		"Staff",
	)
	if err != nil {
		return model.Employee{}, err
	}

	return emp, nil
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
