package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc{
	return func(c echo.Context) error {
		authHeader:=c.Request().Header.Get("Authorization")
		if authHeader==""{
			data:=map[string]interface{}{
				"Message":"Token is missing",
			}
			return c.JSON(http.StatusUnauthorized, data)
		}
		parts:=strings.Split(authHeader, " ")
		if len(parts) !=2 || parts[0] !="Bearer"{
			data:=map[string]interface{}{
				"message":"Invalid Header format",
			}
			return c.JSON(http.StatusUnauthorized, data)

		}
		tokenString:=parts[1]
		token, err:=jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err !=nil{
			data:=map[string]interface{}{
				"Message":"Token validation failed",
			}
			return c.JSON(http.StatusUnauthorized, data)
		}
		claims, ok:=token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid{
			data:=map[string]interface{}{
				"message":"Invalid token claim",
			}
			return c.JSON(http.StatusUnauthorized, data)
		}
		userID, ok:=claims["userID"].(string)
		if !ok{
			data:=map[string]interface{}{
				"Message":"User ID is not found in token claims",
			}
			return c.JSON(http.StatusUnauthorized, data)
		}
		role, ok:=claims["role"].(string)
		if !ok{
			data:=map[string]interface{}{
				"message":"role is not found in token claims",
			}
			return c.JSON(http.StatusUnauthorized, data)
		}
		c.Set("userID", userID)
		c.Set("role", role)

		return next(c)

	}

}