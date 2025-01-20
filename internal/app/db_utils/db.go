package db

import (
	"fmt"

	"github.com/meles-zawude-e/configs"
	"github.com/meles-zawude-e/internal/app/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg configs.DatabaseConfigration) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d name=%s username=%s password=%s", cfg.Host, cfg.Port, cfg.Name, cfg.Username, cfg.Password)
	hospitalDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = hospitalDB
	err = hospitalDB.AutoMigrate(&entities.Ambulance{})
	if err != nil {
		return nil, err
	}

	return hospitalDB, nil
}
