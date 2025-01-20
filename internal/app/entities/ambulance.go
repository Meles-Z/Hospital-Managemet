package entities

//Ambulance: Tracks ambulance availability, usage, and assignments.

type Ambulance struct {
	Model
	License            string
	DriverId           string
	StaffId            string
	AvaliabilityStatus string
	CurrentLocation    string
	PatientId          string
}
