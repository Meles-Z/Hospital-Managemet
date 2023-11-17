package router

import (
	"github.com/labstack/echo/v4"
	"github.com/meles-zawude-e/handler"
	"github.com/meles-zawude-e/middleware"
)

func SetUpRouter(c *echo.Echo) {
	v1 := c.Group("/api")
	v1.POST("/admin", handler.CreateAdmin)
	v1.GET("/admin/:id", handler.GetAdminById)
	v1.DELETE("/admin/:id", handler.DeleteAdmin)
	v1.POST("/login", handler.Login)

	//petient router
	v3 := c.Group("/api")
	v3.POST("/petient", handler.CreatePatient)
	v3.GET("/petient", handler.GetAllPetient)
	v3.GET("/petient/:id", handler.GetPatientByID)
	v3.PUT("/petient/:id", handler.UpdataPetient)
	v3.DELETE("/petient/:id", handler.DeletePetient)

	//doctor router
	v2 := c.Group("/api")
	v2.Use(middleware.ValidateToken)
	v2.POST("/doctor", handler.CreateDoctor)
	v2.GET("/doctor", handler.GetAllDoctor)
	v2.GET("/doctor/:id", handler.GetDoctorById)
	v2.PUT("/doctor/:id", handler.UpdateDoctor)
	v2.DELETE("/doctor/:id", handler.DeleteDoctor)

	//technician router
	v2.POST("/tec", handler.CreateTechnician)
	v2.GET("/tec", handler.GetAllTechnicain)
	v2.GET("/tec/:id", handler.GetTechnicainById)
	v2.PUT("/tec/:id", handler.UpdadeTechnicain)
	v2.DELETE("/tec/:id", handler.DeleteTechnician)

	//Result router
	v2.POST("/res", handler.CreateResult)
	v2.GET("/res", handler.GetAllResult)
	v2.GET("/res/:id", handler.GetResultById)
	v2.PUT("/res/:id", handler.UpdateResult)
	v2.DELETE("/res/:id", handler.DeleteResult)

	//Request router
	v2.POST("/lab", handler.CreateLabRequest)
	v2.GET("/lab", handler.GetAllLabRequest)
	v2.GET("/lab/:id", handler.GetLabRequestById)
	v2.PUT("/lab/:id", handler.UpdateLabRequest)
	v2.DELETE("/lab/:id", handler.DeleteLabRequest)

	//appointment router
	v3.POST("/appt", handler.PatientApointement)
	v3.POST("/appt-e", handler.SendNotificationByEmail)
	v3.GET("/appt", handler.GetAllApointment)
	v3.GET("/appt/:id", handler.GetAppointmentById)
	v3.PUT("/appt/:id", handler.UpdataAppointment)
	v3.DELETE("/appt/:id", handler.DeleteAppointment)

	//paying bill
	v3.POST("/pay", handler.AcceptPayment)
	v3.GET("/pay-v", handler.VerifyPayment)
	v3.POST("/pay-s", handler.SplitPayment)
	v3.GET("/pay-b", handler.GetBank)
	v3.POST("/pay-t", handler.TransferBill)
	//pharmacy
	v2.POST("/ph", handler.CreatePharmatics)
	v2.GET("/ph", handler.GetAllPharmatics)
	v2.GET("/ph/:id", handler.GetAllPharmaticsByID)
	v2.PUT("/ph/:id", handler.UpdatePharmatics)
	v2.DELETE("/ph/:id", handler.DeletePharmatics)

	//receptionist
	v2.POST("/rs", handler.CreateReceptionest)
	v2.GET("/rs", handler.GetAllReceptionist)
	v2.GET("/rs/:id", handler.GetAllReceptionistByID)
	v2.PUT("/rs/:id", handler.UpdateReceptionist)
	v2.DELETE("/rs/:id", handler.DeleteReceptionist)

	//medicine
	v2.POST("/medicine", handler.CreateMedicine)
	v2.GET("/medicine", handler.GetAllMedicine)
	v2.GET("/medicine/:id", handler.GetAllMedicineById)
	v2.PUT("/medicine/:id", handler.UpdateMedicine)
	v2.DELETE("/medicine/:id", handler.DeleteMedicine)
}
