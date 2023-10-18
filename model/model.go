package model

import (
	"time"

	"github.com/google/uuid"
)

type HospAdmin struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Phone     string     `json:"phone"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
}
type HospUser struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Phone    string    `json:"phone"`
	Role     string    `json:"role"`
	Doctors  []Doctor  `json:"doctors" gorm:"foreignKey:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	// Petients   []Petient   `json:"petients" gorm:"foreignKey:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	HospAdmins []HospAdmin `json:"hospAdmins" gorm:"foreignKey:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Technicians []Technician `json:"technicians" gorm:"foreignKey:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

type Petient struct {
	ID                uuid.UUID  `json:"id" gorm:"primaryKey"`
	Name              string     `json:"name"`
	Email             string     `json:"email"`
	Phone             string     `json:"phone"`
	Address           string     `json:"address"`
	MedicalHistory    string     `json:"m_history"`
	InsuranceInfo     string     `json:"I_info"`
	Gender            string     `json:"gender"`
	EmergencyContanct string     `json:"e_contact"` //I will give attetion to this
	CreatedAt         time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt         *time.Time `json:"deleted_at" gorm:"default:null"`
}

type Doctor struct {
	ID             uuid.UUID `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Password       string    `json:"-"`
	Phone          string    `json:"phone"`
	Role           string    `json:"role"`
	Specialization string    `json:"spec"`
	OfficeNumber   string    `json:"officeN"`
	AdminID        uuid.UUID `json:"admin_id"`
	//avaliablity
	CreatedAt time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
}

type Technician struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Phone     string     `json:"phone"`
	AdminID   uuid.UUID  `json:"admin_id"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
}

type LabRequest struct {
	RequestID    uuid.UUID `json:"id" gorm:"primaryKey"`
	TechnicianID uuid.UUID `json:"t-id"`
	DoctorID     uuid.UUID `json:"d-id"`
	PetientID    uuid.UUID `json:"p_id"`
	//priority
	TestingType string `json:"t_type"`
	//status
	Result    string     `json:"result"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
}

type ResultDelivery struct {
	ResultID     uuid.UUID  `json:"id" gorm:"primaryKey"`
	PetientID    uuid.UUID  `json:"p_id"`
	TestType     string     `json:"t-type"`
	ResultStatus string     `json:"r-status"`
	CreatedAt    time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt    *time.Time `json:"deleted_at" gorm:"default:null"`
}

// type PharmacyIntegration struct {
// 	MedicationList     string
// 	PrescriptionId     string
// 	PetientInformation string
// 	DoctorID           uuid.UUID
// 	Status             string
// 	CreatedAt          time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
// 	UpdatedAt          time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
// 	DeletedAt          *time.Time `json:"deleted_at" gorm:"default:null"`
// }
