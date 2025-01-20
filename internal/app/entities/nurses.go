package entities

type Nurse struct {
	Name         string      `json:"name"`
	PhoneNumber  string      `json:"phoneNumber"`
	Email        string      `json:"email"`
	Password     string      `json:"password"`
	DepartmentId string      `json:"departmentId"`
	Department   *Department `gorm:"foreignkey:DepartmentID" json:"department"`
}
