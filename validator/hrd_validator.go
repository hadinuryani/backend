package validator

// validasi tambah data pegawai
type EmployeeRequest struct {
	Name    string `json:"name" binding:"required,min=3,max=100"`
	Address string `json:"address" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
}

// validasi resign 
type UpdateEmployeeStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=active inactive"`
}