package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

// daftar semua route
func RegisterAllRoutes(router *gin.Engine){
	// root
	api := router.Group("/api")
	api.GET("/",func(ctx *gin.Context){
		ctx.JSON(200,gin.H{
			"message" : os.Getenv("APP_NAME"),
			"version" : "1.0",
		})
	})

	// daftar route 
	HRDRoutes(api)

}