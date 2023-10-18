package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/meles-zawude-e/database"
	"github.com/meles-zawude-e/model"
	"golang.org/x/crypto/bcrypt"
)

func CreateTechnician(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "Only admin can create a technician",
		}
		return c.JSON(http.StatusForbidden, data)
	}

	tech := new(model.HospUser)
	if err := c.Bind(tech); err != nil {
		data := map[string]interface{}{
			"message": "Could not bind technician information",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	db := database.GetDB()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(tech.Password), 14)
	if err != nil {
		data := map[string]interface{}{
			"message": "Password is not hashed correctly",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	adminID := c.Get("userID").(string)
	adminUUID, err := uuid.Parse(adminID)
	if err != nil {
		data := map[string]interface{}{
			"message": "ID is not parsed successfully",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	technician := model.HospUser{
		ID:       uuid.New(),
		Name:     tech.Name,
		Email:    tech.Email,
		Password: string(hashedPassword),
		Phone:    tech.Phone,
		Role:     tech.Role,
	}

	if err := db.Create(&technician).Error; err != nil {
		data := map[string]interface{}{
			"message": "Could not create technician",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	technic := model.Technician{
		ID:       technician.ID,
		Name:     tech.Name,
		Email:    tech.Email,
		Password: string(hashedPassword),
		Phone:    tech.Phone,
		AdminID:  adminUUID,
	}

	if err := db.Create(&technic).Error; err != nil {
		data := map[string]interface{}{
			"message": "Could not create technician",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	resp := map[string]interface{}{
		"message": "Technician is created successfully",
		"data":    technic,
	}
	return c.JSON(http.StatusOK, resp)
}

func GetTechnicainById(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "Only admin can create a technician",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	id := c.Param("id")
	db := database.GetDB()
	var technician model.HospUser

	if err := db.Where("id=?", id).First(&technician).Error; err != nil {
		data := map[string]interface{}{
			"message": " technician not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	resp := map[string]interface{}{
		"message": "technician found",
		"data":    technician,
	}
	return c.JSON(http.StatusOK, resp)
}

func GetAllTechnicain(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "Only admin can create a technician",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	db := database.GetDB()
	var tech []model.Technician
	if err := db.Find(&tech).Error; err != nil {
		data := map[string]interface{}{
			"message": "technicians not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	resp := map[string]interface{}{
		"message": "technician found",
		"data":    tech,
	}
	return c.JSON(http.StatusOK, resp)
}

func UpdadeTechnicain(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "Only admin can create a technician",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	id := c.Param("id")
	db := database.GetDB()
	updateTech := new(model.HospUser)
	if err := c.Bind(updateTech); err != nil {
		data := map[string]interface{}{
			"message": " Not bind technician information",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusForbidden, data)
	}
	var technician model.HospUser

	if err := db.Where("id=?", id).First(&technician).Error; err != nil {
		data := map[string]interface{}{
			"message": " technician not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	technician.Name = updateTech.Name
	technician.Email = updateTech.Email
	technician.Phone = updateTech.Phone
	technician.Role = updateTech.Role

	if err := db.Save(&technician).Error; err != nil {
		data := map[string]interface{}{
			"message": " technician information not updated",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	resp := map[string]interface{}{
		"message": "technician information updated successfully",
		"data":    technician,
	}
	return c.JSON(http.StatusOK, resp)

}

func DeleteTechnician(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "Only admin can create a technician",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	id := c.Param("id")
	db := database.GetDB()
	var technician model.HospUser

	if err := db.Where("id=?", id).First(&technician).Error; err != nil {
		data := map[string]interface{}{
			"message": " technician not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	if err := db.Delete(&technician).Error; err != nil {
		data := map[string]interface{}{
			"message": "Could not delete technician",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	resp := map[string]interface{}{
		"message": "technician deleted successfully",
	}
	return c.JSON(http.StatusOK, resp)
}
