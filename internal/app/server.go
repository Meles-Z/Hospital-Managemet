package app

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/meles-zawude-e/configs"
	db "github.com/meles-zawude-e/internal/app/db_utils"
	"github.com/meles-zawude-e/internal/app/repository"
	"github.com/meles-zawude-e/internal/app/services"
	"gorm.io/gorm"
)

// first let me see how server are working // it start with importing service
// what it call in import part // it calls services
// let me start

type IServer interface {
	Start() error
}
type Server struct {
	DB                *gorm.DB
	cfg               configs.Config
	departmentService services.IDepartmentService
	doctorService     services.IDoctorService
}

func NewServer(cfg configs.Config) IServer {
	mainDb, err := db.InitDB(cfg.Database)
	if err != nil {
		log.Fatalf("Error to initalize database:%s", err)
	}

	departmentRepo := repository.NewRepositoryDepartment(mainDb)
	doctorRepo := repository.NewDoctorRepository(mainDb)
	departmentSvc, err := services.NewDepartmentService(departmentRepo)
	if err != nil {
		log.Fatal("Error to init department services", err)
	}
	doctorSvc, err := services.NewDoctorService(doctorRepo)
	if err != nil {
		log.Fatal("Error to init doctor services", err)
	}

	return &Server{
		DB:                mainDb,
		cfg:               cfg,
		departmentService: departmentSvc,
		doctorService:     doctorSvc,
	}

}
func (s *Server) Start() error {
	echo := echo.New()
	Router(echo, *s)
	return echo.Start(fmt.Sprintf("%d", s.cfg.Database.Port))
}
