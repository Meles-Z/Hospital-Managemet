package handler

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/meles-zawude-e/database"
	"github.com/meles-zawude-e/model"
)

func PatientApointement(c echo.Context) error {
	db := database.GetDB()
	app := new(model.Appointment)

	if err := c.Bind(app); err != nil {
		data := map[string]interface{}{
			"message": "data not binded in appointment",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	var doctor model.Doctor
	if err := db.Where("id=?", app.DoctorID).First(&doctor).Error; err != nil {
		data := map[string]interface{}{
			"message": "doctor is not found by this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	var patient model.Petient
	if err := db.Where("id=?", app.PetientID).First(&patient).Error; err != nil {
		data := map[string]interface{}{
			"message": "patient is not found by this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	var lastAppointment model.Appointment
	if err := db.Where("doctor_id = ?", app.DoctorID).Order("end DESC").Limit(1).Find(&lastAppointment).Error; err != nil {
		app.Start = time.Now()
	} else {
		app.Start = lastAppointment.End
	}

	app.End = app.Start.Add(30 * time.Minute)

	appoint := model.Appointment{
		ID:        uuid.New(),
		Email:     app.Email,
		Start:     app.Start,
		End:       app.End,
		DoctorID:  app.DoctorID,
		PetientID: app.PetientID,
	}

	if err := db.Create(&appoint).Error; err != nil {
		data := map[string]interface{}{
			"message": "Appointment is not created",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	resp := map[string]interface{}{
		"message": "Appointment created successfully",
		"data":    appoint,
	}
	return c.JSON(http.StatusOK, resp)
}

func GetAllApointment(c echo.Context) error {
	db := database.GetDB()
	var app []model.Appointment

	if err := db.Find(&app).Error; err != nil {
		data := map[string]interface{}{
			"message": "appointment not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	resp := map[string]interface{}{
		"message": "Appointment found successfully",
		"data":    app,
	}
	return c.JSON(http.StatusOK, resp)
}

func GetAppointmentById(c echo.Context) error {
	db := database.GetDB()
	id := c.Param("id")
	var app model.Appointment

	if err := db.Where("id=?", id).First(&app).Error; err != nil {
		data := map[string]interface{}{
			"message": "appointment not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	resp := map[string]interface{}{
		"message": "appointmet found",
		"data":    app,
	}
	return c.JSON(http.StatusOK, resp)
}

func UpdataAppointment(c echo.Context) error {
	db := database.GetDB()
	id := c.Param("id")

	updatedApp := new(model.Appointment)
	if err := c.Bind(updatedApp); err != nil {
		data := map[string]interface{}{
			"message": "appointment not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}
	var oldAppointment model.Appointment
	if err := db.Where("id=?", id).First(&oldAppointment).Error; err != nil {
		data := map[string]interface{}{
			"message": "appointment not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	oldAppointment.PetientID = updatedApp.PetientID
	oldAppointment.DoctorID = updatedApp.DoctorID
	oldAppointment.Start = updatedApp.Start
	oldAppointment.End = updatedApp.End

	if err := db.Save(&oldAppointment).Error; err != nil {
		data := map[string]interface{}{
			"message": "appointment not updated successully",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	resp := map[string]interface{}{
		"message": "Appointment updated successlly",
		"data":    oldAppointment,
	}
	return c.JSON(http.StatusOK, resp)
}

func DeleteAppointment(c echo.Context) error {
	db := database.GetDB()
	id := c.Param("id")

	var app model.Appointment
	if err := db.Where("id=?", id).First(&app).Error; err != nil {
		data := map[string]interface{}{
			"message": "appointment not found by this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	if err := db.Delete(&app).Error; err != nil {
		data := map[string]interface{}{
			"message": "appointment not deleted",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	resp := map[string]interface{}{
		"message": "Appointment deleted successfully",
	}
	return c.JSON(http.StatusOK, resp)
}
