package model

import (
	"github.com/google/uuid"
)

type Dermantology struct {
	DoctorID    uuid.UUID `json:"doctor_id"`
	PetientsID  uuid.UUID `json:"Patint_id"`
	PatientName string    `json:"Patient_name"`
}

type Optimology struct {
	DoctorID    uuid.UUID `json:"doctor_id"`
	PetientsID  uuid.UUID `json:"Patint_id"`
	PatientName string    `json:"Patient_name"`
}
type Cordiologyy struct {
	DoctorID    uuid.UUID `json:"doctor_id"`
	PetientsID  uuid.UUID `json:"Patint_id"`
	PatientName string    `json:"Patient_name"`
}
