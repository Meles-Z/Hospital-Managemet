package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/meles-zawude-e/database"
	"github.com/meles-zawude-e/model"
	"golang.org/x/crypto/bcrypt"
)

func CreateReceptionest(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "Only admin can create Pharmatics",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	db := database.GetDB()
	r := new(model.HospUser)

	if err := c.Bind(r); err != nil {
		data := map[string]interface{}{
			"message": "not binded successfully",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), 14)
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
	rec_u := model.HospUser{
		ID:       uuid.New(),
		Name:     r.Name,
		Email:    r.Email,
		Password: string(hashedPassword),
		Phone:    r.Phone,
		Role:     r.Role,
	}

	if err := db.Create(&rec_u).Error; err != nil {
		data := map[string]interface{}{
			"message": "Could not create receptionist",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	receptionist := model.Receptionest{
		ID:       rec_u.ID,
		Name:     r.Name,
		Email:    r.Email,
		Password: string(hashedPassword),
		Phone:    r.Phone,
		AdminID:  adminUUID,
	}
	if err := db.Create(&receptionist).Error; err != nil {
		data := map[string]interface{}{
			"message": "Receptionist not created successfully",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	resp := map[string]interface{}{
		"message": "Receptionist created successfully",
		"data":    receptionist,
	}
	return c.JSON(http.StatusOK, resp)
}

func GetAllReceptionist(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "Only admin can check Receptionest",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	db := database.GetDB()
	var receptionist []model.Receptionest
	if err := db.Find(&receptionist).Error; err != nil {
		data := map[string]interface{}{
			"message": "Receptionist not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	resp := map[string]interface{}{
		"message": "Receptionist found successfully",
		"data":    receptionist,
	}
	return c.JSON(http.StatusOK, resp)
}

func GetAllReceptionistByID(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "Only admin can check Receptionst",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	db := database.GetDB()
	id := c.Param("id")
	var receptionist model.Receptionest
	if err := db.Where("id=?", id).First(&receptionist).Error; err != nil {
		data := map[string]interface{}{
			"message": "Receptionist not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	resp := map[string]interface{}{
		"message": "Receptionist found successfully",
		"data":    receptionist,
	}
	return c.JSON(http.StatusOK, resp)
}

// I want to update receptionist by cascade
func UpdateReceptionist(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "Only admin can update Receptions",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	db := database.GetDB()
	id := c.Param("id")

	updateRecep := new(model.Receptionest)
	if err := c.Bind(updateRecep); err != nil {
		data := map[string]interface{}{
			"message": "Not bind correctly",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	var receptionist model.Receptionest
	if err := db.Where("id=?", id).First(&receptionist).Error; err != nil {
		data := map[string]interface{}{
			"message": "Receptionist not found with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	receptionist.Name = updateRecep.Name
	receptionist.Email = updateRecep.Email
	receptionist.Phone = updateRecep.Phone

	if err := db.Save(&receptionist).Error; err != nil {
		data := map[string]interface{}{
			"message": "Not updated successfully",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}
	resp := map[string]interface{}{
		"message": "Receptionist updated successfully",
		"data":    receptionist,
	}
	return c.JSON(http.StatusOK, resp)

}

// I use cascade form to delete Pharmatics
func DeleteReceptionist(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "Only admin can delete Receptionist",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	db := database.GetDB()
	id := c.Param("id")
	var recp model.HospUser
	if err := db.Where("id=?", id).First(&recp).Error; err != nil {
		data := map[string]interface{}{
			"message": "Receptionist not found with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	if err := db.Delete(&recp).Error; err != nil {
		data := map[string]interface{}{
			"message": "Receptionist not found with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	resp := map[string]interface{}{
		"message": "Receptionist deleted successfully",
	}
	return c.JSON(http.StatusOK, resp)

}

// func ReceptionistCreatePatient(c echo.Context) error{
	
// }