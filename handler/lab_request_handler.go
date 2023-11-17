package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/meles-zawude-e/database"
	"github.com/meles-zawude-e/model"
)

func CreateLabRequest(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" && role != "doctor" {
		data := map[string]interface{}{
			"message": "only admin and doctor can create lab request",
		}
		return c.JSON(http.StatusForbidden, data)
	}

	db := database.GetDB()
	labRequest := new(model.LabRequest)
	if err := c.Bind(labRequest); err != nil {
		data := map[string]interface{}{
			"message": "not bind request data",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	var petient model.Petient
	if err := db.Where("id=?", labRequest.PetientID).First(&petient).Error; err != nil {
		data := map[string]interface{}{
			"message": "no patient with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	var doctor model.Doctor
	if err := db.Where("id=?", labRequest.DoctorID).First(&doctor).Error; err != nil {
		data := map[string]interface{}{
			"message": "no doctor with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	var technician model.Technician
	if err := db.Where("id=?", labRequest.TechnicianID).First(&technician).Error; err != nil {
		data := map[string]interface{}{
			"message": "no technician with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	lab := model.LabRequest{
		RequestID:    uuid.New(),
		TechnicianID: labRequest.TechnicianID,
		DoctorID:     labRequest.DoctorID,
		PetientID:    labRequest.PetientID,
		Priority:     labRequest.Priority,
		TestingType:  labRequest.TestingType,
		Status:       labRequest.Status,
	}

	if err := db.Create(&lab).Error; err != nil {
		data := map[string]interface{}{
			"message": "Request is not created",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}
	resp := map[string]interface{}{
		"message": "Request is created successfully",
		"data":    lab,
	}
	return c.JSON(http.StatusOK, resp)

}

func GetLabRequestById(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" && role != "doctor" {
		data := map[string]interface{}{
			"message": "only admin and doctor can create lab request",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	id := c.Param("id")
	db := database.GetDB()
	var request model.LabRequest

	if err := db.Where("request_id=?", id).First(&request).Error; err != nil {
		data := map[string]interface{}{
			"message": "user resut not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	resp := map[string]interface{}{
		"message": "user result is found",
		"data":    request,
	}
	return c.JSON(http.StatusOK, resp)

}
func GetAllLabRequest(c echo.Context) error {

	role := c.Get("role").(string)
	if role != "admin" && role != "doctor" {
		data := map[string]interface{}{
			"message": "only admin and doctor can create lab request",
		}
		return c.JSON(http.StatusForbidden, data)
	}

	db := database.GetDB()
	var request []model.LabRequest

	if err := db.Find(&request).Error; err != nil {
		data := map[string]interface{}{
			"message": "lab request not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	resp := map[string]interface{}{
		"message": "user results found",
		"data":    request,
	}
	return c.JSON(http.StatusOK, resp)

}

func UpdateLabRequest(c echo.Context) error {

	role := c.Get("role").(string)
	if role != "admin" && role != "doctor" {
		data := map[string]interface{}{
			"message": "only admin and doctor can create lab request",
		}
		return c.JSON(http.StatusForbidden, data)
	}

	id := c.Param("id")
	db := database.GetDB()
	updatedResult := new(model.LabRequest)

	if err := c.Bind(updatedResult); err != nil {
		data := map[string]interface{}{
			"message": "result not binded correctly",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	var oldResult model.LabRequest
	if err := db.Where("request_id=?", id).First(&oldResult).Error; err != nil {
		data := map[string]interface{}{
			"message": "request not found by this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	oldResult.TechnicianID = updatedResult.TechnicianID
	oldResult.DoctorID = updatedResult.DoctorID
	oldResult.PetientID = updatedResult.PetientID
	oldResult.Priority = updatedResult.Priority
	oldResult.TestingType = updatedResult.TestingType
	oldResult.Status = updatedResult.Status

	if err := db.Save(&oldResult).Error; err != nil {
		data := map[string]interface{}{
			"message": "result not updated successfully",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	resp := map[string]interface{}{
		"message": "user result updated successfully",
		"data":    oldResult,
	}
	return c.JSON(http.StatusOK, resp)

}

func DeleteLabRequest(c echo.Context) error {

	role := c.Get("role").(string)
	if role != "admin" && role != "doctor" {
		data := map[string]interface{}{
			"message": "only admin and doctor can create lab request",
		}
		return c.JSON(http.StatusForbidden, data)
	}

	id := c.Param("id")
	db := database.GetDB()
	var result model.LabRequest
	if err := db.Where("request_id=?", id).First(&result).Error; err != nil {
		data := map[string]interface{}{
			"message": "request not found with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	if err := db.Delete(&result).Error; err != nil {
		data := map[string]interface{}{
			"message": "lab request not deleted",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	resp := map[string]interface{}{
		"message": "result deleted successfully",
	}
	return c.JSON(http.StatusOK, resp)

}
