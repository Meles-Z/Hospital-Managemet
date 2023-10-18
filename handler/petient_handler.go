package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/meles-zawude-e/database"
	"github.com/meles-zawude-e/model"
)

func CreatePatient(c echo.Context) error {
	db := database.GetDB()
	pet := new(model.Petient)

	if err := c.Bind(pet); err != nil {
		data := map[string]interface{}{
			"message": "not bind successfully",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	petient := model.Petient{
		ID:                uuid.New(),
		Name:              pet.Name,
		Email:             pet.Email,
		Phone:             pet.Phone,
		Address:           pet.Address,
		MedicalHistory:    pet.MedicalHistory,
		InsuranceInfo:     pet.InsuranceInfo,
		Gender:            pet.Gender,
		EmergencyContanct: pet.EmergencyContanct,
	}
	if err := db.Create(&petient).Error; err != nil {
		data := map[string]interface{}{
			"message": "patient is not created successfully",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Petient registered successfully",
		"data":    petient,
	}
	return c.JSON(http.StatusOK, response)
}

func GetPatientByID(c echo.Context) error {
	db := database.GetDB()
	id := c.Param("id")

	var petient model.Petient
	if err := db.Where("id=?", id).First(&petient).Error; err != nil {
		data := map[string]interface{}{
			"message": "petient is not fond with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	response := map[string]interface{}{
		"message": "petient found",
		"data":    petient,
	}
	return c.JSON(http.StatusOK, response)
}

func GetAllPetient(c echo.Context) error{
	db:=database.GetDB()
	var petient []model.Petient

	if err:=db.Find(&petient).Error; err !=nil{
		data:=map[string]interface{}{
			"message":"petient not found",
			"data":err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	response:=map[string]interface{}{
		"message":"petient is foud",
		"data":petient,
	}
	return c.JSON(http.StatusOK, response)
}

func UpdataPetient(c echo.Context) error{
	db:=database.GetDB()
	id:=c.Param("id")

	updatedPetient:=new(model.Petient)
	if err:=c.Bind(updatedPetient); err !=nil{
		data:=map[string]interface{}{
			"message":"data is not binding",
			"data":err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	var petient model.Petient
	if err:=db.Where("id=?", id).First(&petient).Error; err !=nil{
		data:=map[string]interface{}{
			"message":"petient not found by this id",
			"data":err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	petient.Name=updatedPetient.Name
	petient.Email=updatedPetient.Email
	petient.Phone=updatedPetient.Phone
	petient.Address=updatedPetient.Address
	petient.MedicalHistory=updatedPetient.MedicalHistory
	petient.InsuranceInfo=updatedPetient.InsuranceInfo
	petient.Gender=updatedPetient.Gender
	petient.EmergencyContanct=updatedPetient.EmergencyContanct

	if err:=db.Save(&petient).Error; err !=nil{
		data:=map[string]interface{}{
			"message":"could not update petient information",
			"data":err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response:=map[string]interface{}{
		"message":"Petient updated successfully",
		"data":petient,
	}
	return c.JSON(http.StatusOK, response)
}

func DeletePetient(c echo.Context) error{
	id:=c.Param("id")
	db:=database.GetDB()

	var petient model.Petient
	if err:=db.Where("id=?", id).First(&petient).Error; err !=nil{
		data:=map[string]interface{}{
			"message":"petient not found",
			"data":err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	
	if err:=db.Delete(&petient).Error; err !=nil{
		data:=map[string]interface{}{
			"message":"petient not deleted succssfully",
			"data":err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	response:=map[string]interface{}{
		"message":"patient deleted successfully",
		
	}
	return c.JSON(http.StatusOK, response)
}
