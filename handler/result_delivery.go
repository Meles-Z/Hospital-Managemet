package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/meles-zawude-e/database"
	"github.com/meles-zawude-e/model"
)

func CreateResult(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" && role != "technician" {
		data := map[string]interface{}{
			"message": "Only admin and technician can create result",
		}
		return c.JSON(http.StatusForbidden, data)
	}

	db := database.GetDB()
	res := new(model.ResultDelivery)
	if err := c.Bind(res); err != nil {
		data := map[string]interface{}{
			"message": "Data not binding",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	// Check if the provided PetientID exists in the database
	var petient model.Petient
	if err := db.Where("id = ?", res.PetientID).First(&petient).Error; err != nil {
		data := map[string]interface{}{
			"message": "Petient not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	// Create the ResultDelivery with the verified PetientID
	result := model.ResultDelivery{
		ResultID:     uuid.New(),
		PetientID:    res.PetientID,
		TestType:     res.TestType,
		ResultStatus: res.ResultStatus,
	}

	if err := db.Create(&result).Error; err != nil {
		data := map[string]interface{}{
			"message": "Failed to create result",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	resp := map[string]interface{}{
		"message": "Result is created successfully",
		"data":    result,
	}
	return c.JSON(http.StatusOK, resp)
}

func GetResultById(c echo.Context) error {
	id := c.Param("id")
	db := database.GetDB()
	var result model.ResultDelivery

	if err := db.Where("result_id=?", id).First(&result).Error; err != nil {
		data := map[string]interface{}{
			"message": "user resut not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	resp := map[string]interface{}{
		"message": "user result is found",
		"data":    result,
	}
	return c.JSON(http.StatusOK, resp)

}
func GetAllResult(c echo.Context) error {
	db := database.GetDB()
	var result []model.ResultDelivery

	if err := db.Find(&result).Error; err != nil {
		data := map[string]interface{}{
			"message": "results not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	resp := map[string]interface{}{
		"message": "user results found",
		"data":    result,
	}
	return c.JSON(http.StatusOK, resp)

}

func UpdateResult(c echo.Context) error {
	id := c.Param("id")
	db := database.GetDB()
	updatedResult := new(model.ResultDelivery)

	if err := c.Bind(updatedResult); err != nil {
		data := map[string]interface{}{
			"message": "result not binded correctly",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	var oldResult model.ResultDelivery
	if err := db.Where("result_id=?", id).First(&oldResult).Error; err != nil {
		data := map[string]interface{}{
			"message": "result not found by this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	oldResult.PetientID = updatedResult.PetientID
	oldResult.TestType = updatedResult.TestType
	oldResult.ResultStatus = updatedResult.ResultStatus

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

func DeleteResult(c echo.Context) error {
	id := c.Param("id")
	db := database.GetDB()
	var result model.ResultDelivery
	if err := db.Where("result_id=?", id).First(&result).Error; err != nil {
		data := map[string]interface{}{
			"message": "result not found with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	if err := db.Delete(&result).Error; err != nil {
		data := map[string]interface{}{
			"message": "result not deleted",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	resp := map[string]interface{}{
		"message": "result deleted successfully",
	}
	return c.JSON(http.StatusOK, resp)

}
