package entities

type Patient struct {
	Model
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phoneNumber"`
	Gender       string
	Appointments *[]Appointment `gorm:"foreignKey:PatientID" json:"appointments"`
	Billing      *[]Billing     `gorm:"foreignKey:PatientID" json:"billing"` // lazy loading
	// insurance info or date of birth here
}
