package services

import (
	"github.com/meles-zawude-e/internal/app/entities"
	"github.com/meles-zawude-e/internal/app/repository"
	"gorm.io/gorm"
)

type IDepartmentService interface {
	CreateDepartmenet(dep *entities.Department) (*entities.Department, error)
	GetAllDepartment() ([]entities.Department, error)
	FindDepartmentById(id string) (*entities.Department, error)
	UpdateDepartement(dep *entities.Department) (*entities.Department, error)
	DeleteDepartement(id string) error
	WithTrx(*gorm.DB) DepartementService
}

type DepartementService struct {
	depRepo repository.DepartmentRepository
}

func (svc DepartementService) WithTrx(db *gorm.DB) DepartementService {
	svc.depRepo = svc.depRepo.WithTrx(db)
	return svc
}

func NewDepartmentService(depRepo repository.DepartmentRepository) (IDepartmentService, error) {
	return DepartementService{depRepo: depRepo}, nil
}

func (svc DepartementService) CreateDepartmenet(dep *entities.Department) (*entities.Department, error) {
	newDep, err := svc.depRepo.CreateDepartmenet(dep)
	if err != nil {
		return nil, err
	}
	return newDep, nil
}

func (svc DepartementService) GetAllDepartment() ([]entities.Department, error) {
	dep, err := svc.depRepo.GetAllDepartment()
	if err != nil {
		return nil, err
	}
	return dep, nil
}
func (svc DepartementService) FindDepartmentById(id string) (*entities.Department, error) {
	dep, err := svc.depRepo.FindDepartmentById(id)
	if err != nil {
		return nil, err
	}
	return dep, nil
}
func (svc DepartementService) UpdateDepartement(dep *entities.Department) (*entities.Department, error) {
	dep, err := svc.depRepo.UpdateDepartement(dep)
	if err != nil {
		return nil, err
	}
	return dep, err
}

func (svc DepartementService) DeleteDepartement(id string) error {
	err := svc.depRepo.DeleteDepartement(id)
	if err != nil {
		return err
	}
	return nil
}
