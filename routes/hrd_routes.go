package routes

import (
	"backend-rems/handler"

	"github.com/gin-gonic/gin"
)

func HRDRoutes(router *gin.RouterGroup) {
	api := router.Group("/hrd")
	{	
		api.POST("/roles",handler.SetRole)
		api.PUT("/roles/:id",handler.UpdateRole)
		api.DELETE("/roles/:id",handler.DeleteRole)
		api.GET("/roles",handler.GetRole)
		api.POST("/employees",handler.AddEmployees)
		api.GET("/stores",handler.GetStores)
		// api.GET("/employees", handler.GetAllEmployees)
		// api.GET("/employees/:id", handler.GetEmployeeByID)
		// api.PATCH("/employees/:id/status",handler.UpdateEmployeeStatus)
	}
}
