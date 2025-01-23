package entities

type Doctor struct {
	Model
	Name           string        `json:"name"`
	PhoneNumber    string        `json:"phoneNumber"`
	Email          string        `json:"email"`
	Password       string        `json:"password"`
	Specialization string        `json:"specialization"`
	DepartmentId   string        `json:"departmentId"`
	Appointment    []Appointment `gorm:"foreignKey:DoctorID" json:"appointment"`
}
