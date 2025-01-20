package entities

import "time"

type Staff struct {
	Model
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Role        string
	Adress      string
	Department  Department
	Salary      string
	DateOfJoin  time.Time
	Status      string
}
