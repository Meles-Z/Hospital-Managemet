package handler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/meles-zawude-e/database"
	"github.com/meles-zawude-e/model"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	db := database.GetDB()
	login := new(model.HospUser)

	if err := c.Bind(login); err != nil {
		data := map[string]interface{}{
			"Message": "failed to bind data",
			"data":    err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)

	}

	var user model.HospUser
	if err := db.Where("email = ?", login.Email).First(&user).Error; err != nil {
		data := map[string]interface{}{
			"Message": "user not found",
		}
		return c.JSON(http.StatusUnauthorized, data)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		data := map[string]interface{}{
			"Message": "Invalid password",
		}
		return c.JSON(http.StatusUnauthorized, data)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["phone"] = user.Phone
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		data := map[string]interface{}{
			"message": "Failed to generate token",
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	response := map[string]interface{}{
		"message": "Login successfully",
		"data":    tokenString,
	}
	return c.JSON(http.StatusOK, response)
}
