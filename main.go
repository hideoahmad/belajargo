package main

import (
	"belajargo/controllers"
	"belajargo/database"
	"belajargo/models"
	"belajargo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Hubungkan ke database
	database.Connect()

	// Auto-migrate model ke dalam database
	database.DB.AutoMigrate(&models.User{})

	// Inisialisasi controller dengan database dan session
	controllers.InitController(database.DB, r)

	// Load template dan static file
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// Daftar routes
	routes.RegisterRoutes(r)

	// Jalankan server di port 8080
	r.Run(":8080")
}
