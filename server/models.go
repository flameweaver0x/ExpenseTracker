package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

type Transaction struct {
	gorm.Model
	Category  string
	Amount    float64
	UserID    uint
	Timestamp string
}

var db *gorm.DB
var err error

func initializeDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{}, &Transaction{})
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&user)
	c.JSON(http.StatusOK, user)
}

func listUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	db.Delete(&User{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// You can add similar CRUD operations for Transactions as per your requirements.

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Users routes
	r.POST("/users", createUser)
	r.GET("/users", listUsers)
	r.DELETE("/users/:id", deleteUser)

	// Transactions routes
	// Add your transaction endpoints here e.g., r.POST("/transactions", createTransaction)

	return r
}

func main() {
	initializeDB()
	r := setupRouter()
	r.Run(":8080") // listening and serving on localhost:8080
}