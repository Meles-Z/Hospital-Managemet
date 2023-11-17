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
	ID           uuid.UUID      `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Password     string         `json:"-"`
	Phone        string         `json:"phone"`
	Role         string         `json:"role"`
	Doctors      []Doctor       `json:"doctors" gorm:"foreignKey:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	HospAdmins   []HospAdmin    `json:"hospAdmins" gorm:"foreignKey:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Technicians  []Technician   `json:"technicians" gorm:"foreignKey:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Pharmatics   []Pharmatics   `json:"pharmatics" gorm:"foreignKey:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Receptionest []Receptionest `json:"receptionist" gorm:"foreignKey:ID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
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
	Disease           string     `json:"disease"`
	EmergencyContanct string     `json:"e_contact"`
	CreatedAt         time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt         time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt         *time.Time `json:"deleted_at" gorm:"default:null"`
}

type Doctor struct {
	ID             uuid.UUID  `json:"id" gorm:"primaryKey"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	Password       string     `json:"-"`
	Phone          string     `json:"phone"`
	Role           string     `json:"role"`
	Specialization string     `json:"spec"`
	OfficeNumber   string     `json:"officeN"`
	AdminID        uuid.UUID  `json:"admin_id"`
	CreatedAt      time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt      *time.Time `json:"deleted_at" gorm:"default:null"`
}

type DoctorDepartment struct {
	ID             uuid.UUID  `json:"id" gorm:"primaryKey"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	Password       string     `json:"-"`
	Phone          string     `json:"phone"`
	Role           string     `json:"role"`
	Specialization string     `json:"spec"`
	OfficeNumber   string     `json:"officeN"`
	AdminID        uuid.UUID  `json:"admin_id"`
	CreatedAt      time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt      *time.Time `json:"deleted_at" gorm:"default:null"`
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
	RequestID    uuid.UUID  `json:"id" gorm:"primaryKey"`
	TechnicianID uuid.UUID  `json:"t-id"`
	DoctorID     uuid.UUID  `json:"d-id"`
	PetientID    uuid.UUID  `json:"p_id"`
	Priority     string     `json:"priority"`
	TestingType  string     `json:"t_type"`
	Status       string     `json:"status"`
	CreatedAt    time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt    *time.Time `json:"deleted_at" gorm:"default:null"`
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

type Appointment struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey"`
	Email     string     `json:"email"`
	Start     time.Time  `json:"start"`
	End       time.Time  `json:"end"`
	DoctorID  uuid.UUID  `json:"doctor_id"`
	PetientID uuid.UUID  `json:"patient_id"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
}

type MedicalHistory struct {
}

type Pharmatics struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Phone     string     `json:"phone"`
	AdminID   uuid.UUID  `json:"admin_id"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
}

type Medicine struct {
	ID           uuid.UUID  `json:"id" gorm:"primaryKey"`
	Name         string     `json:"name"`
	Description  string     `json:"desc"`
	Manufacturer string     `json:"menuf"`
	ExpiryDate   string     `json:"exp"`
	Dosage       string     `json:"dosage"`
	Price        string     `json:"price"`
	PhamaticsID  uuid.UUID  `json:"pharmatics_id"`
	Quantity     int        `json:"quantity"`
	CreatedAt    time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt    *time.Time `json:"deleted_at" gorm:"default:null"`
}
type Receptionest struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Phone     string     `json:"phone"`
	AdminID   uuid.UUID  `json:"admin_id"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"default:null"`
}
