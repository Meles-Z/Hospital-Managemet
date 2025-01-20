package entities

import "time"

type Model struct {
	ID        string     `gorm:"primary_key" json:"id,omitempty" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	CreatedAt time.Time  `gorm:"not null" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
