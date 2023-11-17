package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/meles-zawude-e/database"
	"github.com/meles-zawude-e/model"
	"golang.org/x/crypto/bcrypt"
)

func CreatePharmatics(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "Only admin can create Pharmatics",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	db := database.GetDB()
	ph := new(model.HospUser)

	if err := c.Bind(ph); err != nil {
		data := map[string]interface{}{
			"message": "not binded successfully",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ph.Password), 14)
	if err != nil {
		data := map[string]interface{}{
			"message": "password is not incripted successfully",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
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
	pharm_u := model.HospUser{
		ID:       uuid.New(),
		Name:     ph.Name,
		Email:    ph.Email,
		Password: string(hashedPassword),
		Phone:    ph.Phone,
		Role:     ph.Role,
	}

	if err := db.Create(&pharm_u).Error; err != nil {
		data := map[string]interface{}{
			"message": "Could not create technician",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	pharmatics := model.Pharmatics{
		ID:       pharm_u.ID,
		Name:     ph.Name,
		Email:    ph.Email,
		Password: string(hashedPassword),
		Phone:    ph.Phone,
		AdminID:  adminUUID,
	}
	if err := db.Create(&pharmatics).Error; err != nil {
		data := map[string]interface{}{
			"message": "Pharmatics not created successfully",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	resp := map[string]interface{}{
		"message": "Pharmatics created successfully",
		"data":    pharmatics,
	}
	return c.JSON(http.StatusOK, resp)
}

func GetAllPharmatics(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "Only admin can check Pharmatics",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	db := database.GetDB()
	var pharmatics []model.Pharmatics
	if err := db.Find(&pharmatics).Error; err != nil {
		data := map[string]interface{}{
			"message": "Pharmatics not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	resp := map[string]interface{}{
		"message": "pharmatics found successfully",
		"data":    pharmatics,
	}
	return c.JSON(http.StatusOK, resp)
}

func GetAllPharmaticsByID(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "Only admin can check Pharmatics",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	db := database.GetDB()
	id := c.Param("id")
	var pharmatic model.Pharmatics
	if err := db.Where("id=?", id).First(&pharmatic).Error; err != nil {
		data := map[string]interface{}{
			"message": "Pharmatics not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	resp := map[string]interface{}{
		"message": "pharmatic found successfully",
		"data":    pharmatic,
	}
	return c.JSON(http.StatusOK, resp)
}

// I want to update pharmacist by cascade
func UpdatePharmatics(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "Only admin can update Pharmatics",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	db := database.GetDB()
	id := c.Param("id")

	updatePhamatics := new(model.Pharmatics)
	if err := c.Bind(updatePhamatics); err != nil {
		data := map[string]interface{}{
			"message": "Not bind correctly",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	var pharmacist model.Pharmatics
	if err := db.Where("id=?", id).First(&pharmacist).Error; err != nil {
		data := map[string]interface{}{
			"message": "Pharmacist not found with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	pharmacist.Name = updatePhamatics.Name
	pharmacist.Email = updatePhamatics.Email
	pharmacist.Phone = updatePhamatics.Phone

	if err := db.Save(&pharmacist).Error; err != nil {
		data := map[string]interface{}{
			"message": "Not updated successfully",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}
	resp := map[string]interface{}{
		"message": "Pharmatics updated successfully",
		"data":    pharmacist,
	}
	return c.JSON(http.StatusOK, resp)

}

// I use cascade form to delete pharmatics
func DeletePharmatics(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "Only admin can delete Pharmatics",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	db := database.GetDB()
	id := c.Param("id")
	var pharmacist model.HospUser
	if err := db.Where("id=?", id).First(&pharmacist).Error; err != nil {
		data := map[string]interface{}{
			"message": "Pharmacist not found with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	if err := db.Delete(&pharmacist).Error; err != nil {
		data := map[string]interface{}{
			"message": "Pharmacist not found with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	resp := map[string]interface{}{
		"message": "Pharmatics deleted successfully",
	}
	return c.JSON(http.StatusOK, resp)

}
