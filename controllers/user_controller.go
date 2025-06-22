package controllers

import (
	"belajargo/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitController dipanggil dari main.go untuk mengatur koneksi DB & middleware session
func InitController(db *gorm.DB, r *gin.Engine) {
	DB = db
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
}

// Halaman Login
func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// Halaman Register
func ShowRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

// Proses Registrasi
func RegisterUser(c *gin.Context) {
	var user models.User

	// Ambil data dari form
	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	user.Phone = c.PostForm("phone")
	user.Address = c.PostForm("address")
	user.Education = c.PostForm("education")

	// Konversi umur
	age, err := strconv.Atoi(c.PostForm("age"))
	if err != nil {
		c.String(http.StatusBadRequest, "Umur tidak valid")
		return
	}
	user.Age = age

	// Hash password
	password := c.PostForm("password")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.String(http.StatusInternalServerError, "Gagal hash password")
		return
	}
	user.Password = string(hash)

	// Simpan ke database
	if err := DB.Create(&user).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Gagal simpan user: %v", err))
		return
	}

	c.Redirect(http.StatusFound, "/login")
}

// Proses Login
func LoginUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user models.User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.String(http.StatusUnauthorized, "Email tidak ditemukan")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.String(http.StatusUnauthorized, "Password salah")
		return
	}

	// Simpan session user ID
	sess := sessions.Default(c)
	sess.Set("user_id", user.ID)
	sess.Save()

	c.Redirect(http.StatusFound, "/dashboard")
}

// Halaman Dashboard
func ShowDashboard(c *gin.Context) {
	sess := sessions.Default(c)
	userID := sess.Get("user_id")
	if userID == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	var user models.User
	if err := DB.First(&user, userID).Error; err != nil {
		c.String(http.StatusInternalServerError, "Gagal ambil data user")
		return
	}

	c.HTML(http.StatusOK, "dashboard.html", user)
}

// Logout
func Logout(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Clear()
	sess.Save()

	c.Redirect(http.StatusFound, "/login")
}
