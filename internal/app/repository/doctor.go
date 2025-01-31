package repository

import (
	"fmt"

	"github.com/meles-zawude-e/internal/app/entities"
	"gorm.io/gorm"
)

type DoctorRepository interface {
	CreateDoctor(*entities.Doctor) (*entities.Doctor, error)
	GetDoctor() ([]entities.Doctor, error)
	GetAllDoctorById(id string) (*entities.Doctor, error)
	UpdateDoctor(*entities.Doctor) (*entities.Doctor, error)
	DeleteDoctor(id string) error
	WithTrx(*gorm.DB) DoctorRepository
}

type doctorRepoImp struct {
	DB *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) DoctorRepository {
	return &doctorRepoImp{DB: db}
}

func (repo doctorRepoImp) WithTrx(db *gorm.DB) DoctorRepository {
	if db == nil {
		fmt.Println("trx not found")
		return repo
	}
	return repo
}

func (repo doctorRepoImp) CreateDoctor(doctor *entities.Doctor) (*entities.Doctor, error) {
	err := repo.DB.Create(&doctor).Error
	if err != nil {
		return nil, err
	}
	return doctor, nil
}

func (repo doctorRepoImp) GetDoctor() ([]entities.Doctor, error) {
	var doctor []entities.Doctor
	err := repo.DB.Find(&doctor).Error
	if err != nil {
		return nil, err

	}
	return doctor, nil
}
func (repo doctorRepoImp) GetAllDoctorById(id string) (*entities.Doctor, error) {
	var doctor entities.Doctor
	err := repo.DB.Where("id=?", id).Take(&doctor).Error
	if err != nil {
		return nil, err
	}
	return &doctor, err
}

func (repo doctorRepoImp) UpdateDoctor(doctor *entities.Doctor) (*entities.Doctor, error) {
	var existingDoctor entities.Doctor
	err := repo.DB.Model(&existingDoctor).Where("id=?", doctor.ID).Updates(&entities.Doctor{
		Name:           doctor.Name,
		PhoneNumber:    doctor.PhoneNumber,
		Email:          doctor.Email,
		Password:       doctor.Password,
		Specialization: doctor.Specialization,
	}).Scan(&existingDoctor).Error
	if err != nil {
		return nil, err
	}
	return repo.GetAllDoctorById(doctor.ID)
}

func (repo doctorRepoImp) DeleteDoctor(id string) error {
	var doctor entities.Doctor
	err := repo.DB.Where("id=?", id).Delete(&doctor).Error
	if err != nil {
		return err
	}
	fmt.Printf("Doctor %s id deleted!", doctor.Name)
	return nil
}
