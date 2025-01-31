package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/meles-zawude-e/internal/app/entities"
	"github.com/meles-zawude-e/internal/app/services"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Reason string `json:"reason"`
}

func CreateDoctor(departmentSvc services.IDepartmentService) echo.HandlerFunc {
	return func(c echo.Context) error {
		department := new(entities.Department)
		if err := c.Bind(&department); err != nil {
			return c.JSON(http.StatusBadRequest, &ErrorResponse{
				Reason: fmt.Sprintf("Invalid request body:%s", err.Error()),
			})
		}
		trx := c.Get("trx_db").(*gorm.DB)
		dep := departmentSvc.WithTrx(trx)
		
		// create department
		deparm, err := dep.CreateDepartmenet(department)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &ErrorResponse{
				Reason: fmt.Sprintf("Error to create department:%s", err.Error()),
			})
		}
		return c.JSON(http.StatusCreated, deparm)
	}
}
