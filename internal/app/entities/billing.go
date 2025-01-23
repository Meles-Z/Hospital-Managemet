package entities

import (
	"time"

	"github.com/shopspring/decimal"
)

//Payment: Manages payment transactions, modes, and statuses.

type Billing struct {
	Model
	PatientId     string
	Amount        decimal.Decimal
	PaymentStatus string
	BillingDate   time.Time
	// patient left here
}
