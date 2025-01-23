package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

// lazy loading concept here
type Staff struct {
	Model
	FirstName   string          `json:"firstName"`
	LastName    string          `json:"lastName"`
	Email       string          `json:"email"`
	PhoneNumber string          `json:"phoneNumber"`
	Role        string          `json:"role"`
	Address     string          `json:"address"`
	Salary      decimal.Decimal `json:"salary"`
	DateOfJoin  time.Time       `json:"dateOfJoin"`
	Status      string          `json:"status"`
	Department  *[]Department   `json:"department"`
}
