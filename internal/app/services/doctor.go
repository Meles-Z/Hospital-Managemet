package services

import (
	"fmt"

	"github.com/meles-zawude-e/internal/app/entities"
	"github.com/meles-zawude-e/internal/app/repository"
	"gorm.io/gorm"
)

type IDoctorService interface {
	CreateDoctor(*entities.Doctor) (*entities.Doctor, error)
	GetDoctor() ([]entities.Doctor, error)
	GetDoctorById(id string) (*entities.Doctor, error)
	UpdateDoctor(*entities.Doctor) (*entities.Doctor, error)
	DeleteDoctor(id string) error
	WithTrx(*gorm.DB) DoctorService
}

type DoctorService struct {
	doctorRepo repository.DoctorRepository
}

func (svc DoctorService) WithTrx(db *gorm.DB) DoctorService {
	if db == nil {
		fmt.Println("trx is not found")
		return svc
	}
	svc.doctorRepo = svc.doctorRepo.WithTrx(db)
	return svc
}

func NewDoctorService(doctorRepo repository.DoctorRepository) (IDoctorService, error) {
	return DoctorService{doctorRepo: doctorRepo}, nil
}

func (svc DoctorService) CreateDoctor(doctor *entities.Doctor) (*entities.Doctor, error) {
	newDoctor, err := svc.doctorRepo.CreateDoctor(doctor)
	if err != nil {
		return nil, err
	}
	return newDoctor, nil
}

func (svc DoctorService) GetDoctor() ([]entities.Doctor, error) {
	newDoctor, err := svc.doctorRepo.GetDoctor()
	if err != nil {
		return nil, err
	}
	return newDoctor, nil
}

func (svc DoctorService) GetDoctorById(id string) (*entities.Doctor, error) {
	doctor, err := svc.doctorRepo.GetDoctorById(id)
	if err != nil {
		return nil, err
	}
	return doctor, nil
}

func (svc DoctorService) UpdateDoctor(doctor *entities.Doctor) (*entities.Doctor, error) {
	newDoctor, err := svc.doctorRepo.UpdateDoctor(doctor)
	if err != nil {
		return nil, err
	}
	return newDoctor, nil
}

func (svc DoctorService) DeleteDoctor(id string) error {
	err := svc.doctorRepo.DeleteDoctor(id)
	if err != nil {
		return err
	}
	return nil
}
