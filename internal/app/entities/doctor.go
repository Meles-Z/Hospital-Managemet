package entities

type Doctor struct {
	Model
	Name           string         `json:"name"`
	PhoneNumber    string         `json:"phoneNumber"`
	Email          string         `json:"email"`
	Password       string         `json:"password"`
	Specialization string         `json:"specialization"`
	DepartmentId   string         `json:"departmentId"`
	Department     *Department    `gorm:"foreignkey:DepartmentID" json:"department"`
	Appointment    *[]Appointment `gorm:"foreignKey:DoctorID" json:"appointment"` // lazy loading. what it means we not
	// actually want to load appointemnt to doctor directly untill we explicitly called it.
	// schudule
}
