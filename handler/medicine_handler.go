package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/meles-zawude-e/database"
	"github.com/meles-zawude-e/model"
)

func CreateMedicine(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "pharmatics" && role != "admin" {
		data := map[string]interface{}{
			"message": "only pharmatics and admins can create medicine",
		}
		return c.JSON(http.StatusBadRequest, data)
	}
	db := database.GetDB()
	med := new(model.Medicine)
	if err := c.Bind(med); err != nil {
		data := map[string]interface{}{
			"message": "data is not binded ",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	pharmaticsID := c.Get("userID").(string)
	pharmaticsUUID, err := uuid.Parse(pharmaticsID)
	if err != nil {
		data := map[string]interface{}{
			"message": "user id not compatible",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	medicine := model.Medicine{
		ID:           uuid.New(),
		Name:         med.Name,
		Description:  med.Description,
		Manufacturer: med.Manufacturer,
		ExpiryDate:   med.ExpiryDate,
		Dosage:       med.Dosage,
		Price:        med.Price,
		Quantity:     med.Quantity,
		PhamaticsID:  pharmaticsUUID,
	}

	if err := db.Create(&medicine).Error; err != nil {
		data := map[string]interface{}{
			"message": "medicine not created successfully",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	resp := map[string]interface{}{
		"message": "Medicine created successfully",
		"data":    medicine,
	}
	return c.JSON(http.StatusOK, resp)
}

func GetAllMedicine(c echo.Context) error {
	db := database.GetDB()
	var medicine []model.Medicine

	if err := db.Find(&medicine).Error; err != nil {
		data := map[string]interface{}{
			"message": "medicine not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	resp := map[string]interface{}{
		"message": "Medical found",
		"data":    medicine,
	}
	return c.JSON(http.StatusOK, resp)
}

func GetAllMedicineById(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "pharmatics" && role != "admin" {
		data := map[string]interface{}{
			"message": "only pharmatics and admins can create medicine",
		}
		return c.JSON(http.StatusBadRequest, data)
	}
	db := database.GetDB()
	id := c.Param("id")
	var medicine model.Medicine

	if err := db.Where("id=?", id).First(&medicine).Error; err != nil {
		data := map[string]interface{}{
			"message": "medicine not found with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	resp := map[string]interface{}{
		"message": "Medical found",
		"data":    medicine,
	}
	return c.JSON(http.StatusOK, resp)
}

func UpdateMedicine(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "pharmatics" && role != "admin" {
		data := map[string]interface{}{
			"message": "only pharmatics and admins can create medicine",
		}
		return c.JSON(http.StatusBadRequest, data)
	}
	db := database.GetDB()
	id := c.Param("id")
	medc := new(model.Medicine)
	if err := c.Bind(medc); err != nil {
		data := map[string]interface{}{
			"message": "data is not binded ",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	var medicine model.Medicine
	if err := db.Where("id=?", id).First(&medicine).Error; err != nil {
		data := map[string]interface{}{
			"message": "medicine not found with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	pharmaticsID := c.Get("userID").(string)
	pharmaticsUUID, err := uuid.Parse(pharmaticsID)
	if err != nil {
		data := map[string]interface{}{
			"message": "user id not compatible",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	medicine.Name = medc.Name
	medicine.Description = medc.Description
	medicine.Manufacturer = medc.Manufacturer
	medicine.ExpiryDate = medc.ExpiryDate
	medicine.Dosage = medc.Dosage
	medicine.Price = medc.Price
	medicine.Quantity=medc.Quantity
	medicine.PhamaticsID=pharmaticsUUID

	if err := db.Save(&medicine).Error; err != nil {
		data := map[string]interface{}{
			"message": "medicine not found with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	resp := map[string]interface{}{
		"message": "Medical updated successfully",
		"data":    medicine,
	}
	return c.JSON(http.StatusOK, resp)

}

func DeleteMedicine(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "pharmatics" && role != "admin" {
		data := map[string]interface{}{
			"message": "only pharmatics and admins can create medicine",
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	db := database.GetDB()
	id := c.Param("id")
	var medicine model.Medicine
	if err := db.Where("id=?", id).First(&medicine).Error; err != nil {
		data := map[string]interface{}{
			"message": "medicine not found with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	if err := db.Delete(&medicine).Error; err != nil {
		data := map[string]interface{}{
			"message": "medicine not deleted successfully",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	response := map[string]interface{}{
		"message": "Medicine deleted successfully",
	}
	return c.JSON(http.StatusOK, response)

}
