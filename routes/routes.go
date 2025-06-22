package routes

import (
	"belajargo/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Arahkan root "/" ke halaman login
	r.GET("/", controllers.ShowLoginPage)

	r.GET("/login", controllers.ShowLoginPage)
	r.POST("/login", controllers.LoginUser)

	r.GET("/register", controllers.ShowRegisterPage)
	r.POST("/register", controllers.RegisterUser)

	r.GET("/dashboard", controllers.ShowDashboard)
	r.POST("/logout", controllers.Logout)
}

