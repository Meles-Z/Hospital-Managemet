package entities

//Department: Organizes staff and services into departments like Cardiology, Neurology, etc.

type Department struct {
	Model
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Desctiption string   `json:"description"`
	PhoneNumber string   `json:"phoneNumber"`
	Location    string   `json:"location"`
	Doctor      []Doctor `gorm:"foreignKey:DepartmentID" json:"doctor"`
	Nurse       []Nurse  `gorm:"foreignKey:DepartmentID" json:"nurse"`
}
