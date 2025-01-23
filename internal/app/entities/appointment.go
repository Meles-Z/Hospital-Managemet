package entities

import "time"

type Appointment struct {
	Model
	PetientId string    `json:"appointmentID"`
	Date      time.Time `json:"date"`
	Reason    string    `json:"reason"`
	Notes     string    `json:"notes"`
	DoctorId  string    `json:"doctorID"`
	Status    string    `json:"status"`
	Patient   *Patient  `gorm:"foreignKey:PatientID" json:"patientID"`
}
