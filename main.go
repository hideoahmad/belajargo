package main

import (
	"belajargo/controllers"
	"belajargo/database"
	"belajargo/models"
	"belajargo/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file (jika ada, agar bisa digunakan secara lokal)
	_ = godotenv.Load()

	// Buat router
	r := gin.Default()

	// Hubungkan ke database
	database.Connect()

	// Auto-migrate model ke dalam database
	database.DB.AutoMigrate(&models.User{})

	// Inisialisasi controller dengan DB dan session
	controllers.InitController(database.DB, r)

	// Load template HTML dan folder static
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	// Daftarkan semua route
	routes.RegisterRoutes(r)

	// Jalankan server (gunakan PORT dari environment jika ada)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback lokal
	}

	r.Run(fmt.Sprintf(":%s", port))
}
