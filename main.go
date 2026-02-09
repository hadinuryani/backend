package main

import (
	"backend-rems/config"
	"backend-rems/middleware"
	"backend-rems/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	// load env
	if err:= godotenv.Load();err != nil{
		log.Fatal("Error load .env file")
	}

	// koneksi database
	config.ConnectDB()
	defer config.DB.Close()

	// gin router
	r := gin.Default()

	// middleware
	r.Use(middleware.CROSMiddleware())
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.ErrorHandling())

	// data routes
	routes.RegisterAllRoutes(r)

	// start server
	port := os.Getenv("SERVER_PORT")
	log.Printf("server menyala di port %s",port)
	if err := r.Run();err != nil {
		log.Fatalf("gagal menyalakan server %v\n",err)
	}

		
}