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

}
