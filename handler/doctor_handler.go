package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/meles-zawude-e/database"
	"github.com/meles-zawude-e/model"
	"golang.org/x/crypto/bcrypt"
)

func CreateDoctor(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "Only admin can create doctor",
		}
		return c.JSON(http.StatusForbidden, data)
	}

	db := database.GetDB()
	a := new(model.Doctor)

	if err := c.Bind(a); err != nil {
		data := map[string]interface{}{
			"message": "not bind data",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), 14)
	if err != nil {
		data := map[string]interface{}{
			"message": "password is not hashed correctyly",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	d_user := model.HospUser{
		ID:       uuid.New(),
		Name:     a.Name,
		Email:    a.Email,
		Password: string(hashedPassword),
		Phone:    a.Phone,
		Role:     a.Role,
	}

	if err := db.Create(&d_user).Error; err != nil {
		data := map[string]interface{}{
			"message": "user not created in user doctor",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	adminID := c.Get("userID").(string)
	adminUUID, err := uuid.Parse(adminID)
	if err != nil {
		data := map[string]interface{}{
			"message": "Id is not parsed correctly",
		}
		return c.JSON(http.StatusBadRequest, data)
	}
	doctor := model.Doctor{
		ID:             d_user.ID,
		Name:           a.Name,
		Email:          a.Email,
		Password:       string(hashedPassword),
		Phone:          a.Phone,
		Role:           a.Role,
		Specialization: a.Specialization,
		OfficeNumber:   a.OfficeNumber,
		AdminID:        adminUUID,
	}
	if err := db.Create(&doctor).Error; err != nil {
		data := map[string]interface{}{
			"message": "user not created in user doctor",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "doctor is created successfully",
		"data":    doctor,
	}
	return c.JSON(http.StatusOK, response)

}

func GetAllDoctor(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "only admin can see all doctors",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	db := database.GetDB()
	var users []model.Doctor

	if err := db.Find(&users).Error; err != nil {
		data := map[string]interface{}{
			"message": "doctor not found",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	return c.JSON(http.StatusOK, users)
}

func GetDoctorById(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "only admin can see all doctors",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	id := c.Param("id")
	db := database.GetDB()
	var doctor model.Doctor

	if err := db.Where("id = ?", id).First(&doctor).Error; err != nil {
		data := map[string]interface{}{
			"message": "Error fetching doctor",
			"error":   err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	return c.JSON(http.StatusOK, doctor)
}

 func UpdateDoctor(c echo.Context) error{
	role:=c.Get("role").(string)
	if role != "admin"{
		data:=map[string]interface{}{
			"message":"only admin can update the information",
		}
		return c.JSON(http.StatusForbidden, data)
	}

	db:=database.GetDB()
	id:=c.Param("id")
	updatedDoctor:=new(model.Doctor)
	if err:=c.Bind(updatedDoctor); err !=nil{
		data:=map[string]interface{}{
			"message":"new doctor is failed to bind",
			"data":err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)

	}
	var oldInfo model.Doctor
	if err:=db.Where("id=?", id).First(&oldInfo).Error; err !=nil {
		data:=map[string]interface{}{
			"message":"doctor is not found with this id",
			"data":err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
    
	adminID:=c.Get("userID").(string)
	adminUUID, err:=uuid.Parse(adminID)
	if err !=nil{
		data:=map[string]interface{}{
			"message":"doctor id is not parsed correctly",
			
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	oldInfo.Name=updatedDoctor.Name
	oldInfo.Email=updatedDoctor.Email
	oldInfo.Password=updatedDoctor.Password
	oldInfo.Phone=updatedDoctor.Phone
	oldInfo.Specialization=updatedDoctor.Specialization
	oldInfo.OfficeNumber=updatedDoctor.OfficeNumber
	oldInfo.AdminID=adminUUID

	if err:=db.Save(&oldInfo).Error; err !=nil{
		data:=map[string]interface{}{
			"message":"doctor not updated successfully",
			"data":err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	response:=map[string]interface{}{
		"message":"Doctor updated successfully",
		"data":oldInfo,
	}
	return c.JSON(http.StatusOK, response)

 }

func DeleteDoctor(c echo.Context) error {
	role := c.Get("role").(string)
	if role != "admin" {
		data := map[string]interface{}{
			"message": "only admin can see all doctors",
		}
		return c.JSON(http.StatusForbidden, data)
	}
	id := c.Param("id")
	db := database.GetDB()
	var doctor_u model.HospUser
	if err := db.Where("id=?", id).First(&doctor_u).Error; err != nil {
		data := map[string]interface{}{
			"message": "doctor not found with this id",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	if err := db.Delete(&doctor_u).Error; err != nil {
		data := map[string]interface{}{
			"message": "record is not deleted because of some error",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	response := map[string]interface{}{
		"message": "Doctor deleted successfully",
	}
	return c.JSON(http.StatusOK, response)
}
