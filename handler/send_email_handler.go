package handler

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/meles-zawude-e/database"
	"github.com/meles-zawude-e/model"
	mail "github.com/xhit/go-simple-mail/v2"
)

func SendNotificationByEmail(c echo.Context) error {
	db := database.GetDB()

	var appointment model.Appointment
	if err := c.Bind(&appointment); err != nil {
		data := map[string]interface{}{
			"message": "Data not bound to the appointment",
			"error":   err.Error(),
		}
		return c.JSON(http.StatusBadRequest, data)
	}

	if err := db.Where("id = ?", appointment.ID).First(&appointment).Error; err != nil {
		data := map[string]interface{}{
			"message": "Appointment not found in the database",
		}
		return c.JSON(http.StatusNotFound, data)
	}

	// Create an instance of the SMTP client
	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	smtpUsername := "meles.zawdie@gmail.com"
	smtpPassword := "ohmmsappkpbscuzg"

	email := mail.NewMSG()
	email.SetFrom("meles.zawdie@gmail.com").
		AddTo(appointment.Email).
		SetSubject("Appointment Information").
		SetBody(mail.TextPlain, "Hello, Thank you for apply on our hostpitals. \n "+
			"Your appointment details are:\n "+
			"Start Time: "+appointment.Start.Format(time.RFC3339)+"\n"+
			"End Time: "+appointment.End.Format(time.RFC3339)+"\n"+
			"Doctor ID: "+appointment.DoctorID.String()+"\n"+
			"Your ID: "+appointment.PetientID.String())

	// Create a new SMTP client
	client := mail.NewSMTPClient()
	client.Host = smtpHost
	client.Port = smtpPort
	client.Username = smtpUsername
	client.Password = smtpPassword
	client.Encryption = mail.EncryptionSTARTTLS

	client.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	// Connect to the SMTP server
	server, err := client.Connect()
	if err != nil {
		data := map[string]interface{}{
			"message": "Failed to connect to the SMTP server",
			"error":   err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}
	defer server.Close()

	// Send the email
	if err := email.Send(server); err != nil {
		data := map[string]interface{}{
			"message": "Failed to send the email",
			"error":   err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	data := map[string]interface{}{
		"message": "Email sent successfully",
	}
	return c.JSON(http.StatusOK, data)
}
