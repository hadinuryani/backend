package routes

import (
	"backend-rems/handler"

	"github.com/gin-gonic/gin"
)

func HRDRoutes(router *gin.RouterGroup) {
	api := router.Group("/hrd")
	{
		api.GET("/employees", handler.GetAllEmployees)
		api.GET("/employees/:id", handler.GetEmployeeByID)
		api.POST("/employees",handler.AddEmployees)
		api.PATCH("/employees/:id/status",handler.UpdateEmployeeStatus)
	}
}
