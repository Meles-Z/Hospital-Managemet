package repository

import (
	"fmt"

	"github.com/meles-zawude-e/internal/app/entities"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	CreateDepartmenet(dep *entities.Department) (*entities.Department, error)
	GetAllDepartment() ([]entities.Department, error)
	FindDepartmentById(id string) (*entities.Department, error)
	UpdateDepartement(dep *entities.Department) (*entities.Department, error)
	DeleteDepartement(id string) error
	WithTrx(*gorm.DB) DepartmentRepository
}

type departmenentRepImp struct {
	DB *gorm.DB
}

func NewRepositoryDepartment(db *gorm.DB) DepartmentRepository {
	return &departmenentRepImp{DB: db}
}

func (repo departmenentRepImp) WithTrx(trxHandler *gorm.DB) DepartmentRepository {
	if trxHandler == nil {
		fmt.Println("Transaction not found")
		return repo
	}
	repo.DB = trxHandler
	return repo
}

func (repo departmenentRepImp) CreateDepartmenet(dep *entities.Department) (*entities.Department, error) {
	err := repo.DB.Create(&dep).Error
	if err != nil {
		return nil, err
	}
	return dep, nil
}

func (repo departmenentRepImp) GetAllDepartment() ([]entities.Department, error) {
	var dep []entities.Department
	err := repo.DB.Find(&dep).Error
	if err != nil {
		return nil, err
	}
	return dep, nil
}

func (repo departmenentRepImp) FindDepartmentById(id string) (*entities.Department, error) {
	var dep entities.Department
	err := repo.DB.Take(&dep).Error
	if err != nil {
		return nil, err
	}
	return &dep, nil
}
func (repo departmenentRepImp) UpdateDepartement(dep *entities.Department) (*entities.Department, error) {
	var existingDepartment entities.Department
	err := repo.DB.Model(&existingDepartment).First("id=?", dep.ID).Updates(entities.Department{
		Name:        dep.Name,
		Email:       dep.Email,
		Desctiption: dep.Desctiption,
		PhoneNumber: dep.PhoneNumber,
		Location:    dep.Location,
	}).Scan(&existingDepartment).Error
	if err != nil {
		return nil, err
	}
	updatedValue, err := repo.FindDepartmentById(existingDepartment.ID)
	if err != nil {
		return nil, err
	}
	return updatedValue, nil
}
func (repo departmenentRepImp) DeleteDepartement(id string) error {
	var dep entities.Department
	err := repo.DB.Find("id=?", id).Delete(&dep).Error
	if err != nil {
		return err
	}
	return nil
}
