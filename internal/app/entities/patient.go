package entities

type Patient struct {
	Model
	Name         string        `json:"name"`
	Age          int           `json:"age"`
	Email        string        `json:"email"`
	PhoneNumber  string        `json:"phoneNumber"`
	Appointments []Appointment `gorm:"foreignKey:PatientID" json:"appointments"`
}
