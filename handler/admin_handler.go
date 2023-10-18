package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/meles-zawude-e/database"
	"github.com/meles-zawude-e/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateAdmin(c echo.Context) error {
	db := database.GetDB()
	a := new(model.HospUser)

	if err := c.Bind(a); err != nil {
		data := map[string]interface{}{
			"Message": "Error to bind admins",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(a.Password), 14)
	if err != nil {
		data := map[string]interface{}{
			"Message": "failed to generate password",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	u_admin := model.HospUser{
		ID:       uuid.New(),
		Name:     a.Name,
		Email:    a.Email,
		Password: string(hashedPass),
		Phone:    a.Phone,
		Role:     a.Role,
	}
	if err := db.Create(&u_admin).Error; err != nil {
		data := map[string]interface{}{
			"Message": "Error to create hospital user in admin page",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	admin := model.HospAdmin{
		ID:       u_admin.ID,
		Name:     a.Name,
		Email:    a.Email,
		Password: string(hashedPass),
		Phone:    a.Phone,
	}
	if err := db.Create(&admin).Error; err != nil {
		data := map[string]interface{}{
			"Message": "failed to create admin",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	response := map[string]interface{}{
		"Message": "Admin Created Successfully",
		"data":    admin,
	}
	return c.JSON(http.StatusOK, response)

}

func GetAdminById(c echo.Context) error{
	db:=database.GetDB()
	id:=c.Param("id")
    var admin model.HospAdmin
	if err:=db.Where("id=?", id).First(&admin).Error; err !=nil{
		data:=map[string]interface{}{
			"message":"admin not found by this id",
			"data":err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}
	response:=map[string]interface{}{
		"message ":"admin found ",
		"data":admin,
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteAdmin(c echo.Context) error {
	db := database.GetDB()
	id := c.Param("id")

	var h_user model.HospUser

	if err := db.Where("id=?", id).First(&h_user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			data := map[string]interface{}{
				"message": "recourd not found in user table ",
				"data":    err.Error(),
			}
			return c.JSON(http.StatusNotFound, data)
		}
		data := map[string]interface{}{
			"message": "failed to find in the user table in user table",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	if err := db.Delete(&h_user).Error; err != nil {
		data := map[string]interface{}{
			"message": "not deleted from user table",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Record is deleted successfully",
	}
	return c.JSON(http.StatusOK, response)
}
