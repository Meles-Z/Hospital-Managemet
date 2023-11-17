package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/meles-zawude-e/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var gormDB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro to load env file")
	}

	DB_host := os.Getenv("DB_HOST")
	DB_user := os.Getenv("DB_USER")
	DB_name := os.Getenv("DB_NAME")
	DB_password := os.Getenv("DB_PASSWORD")
	DB_port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", DB_user, DB_password, DB_name, DB_host, DB_port)
	gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to database")
}

func GetDB() *gorm.DB {
	return gormDB
}

func AutoMigrateDB() {
	err := gormDB.AutoMigrate(&model.HospAdmin{}, &model.HospUser{}, &model.Petient{}, &model.Doctor{}, &model.Technician{},
		&model.LabRequest{}, &model.ResultDelivery{}, &model.Appointment{}, &model.Pharmatics{}, &model.Receptionest{},
		&model.Medicine{})

	if err != nil {
		log.Fatal(err)
	}
}
