package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("Database connection successfully established")

	db.AutoMigrate(&User{}, &Transaction{})
	log.Println("Database migration completed")
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&user)
	log.Printf("User created: %s\n", user.Name)
	c.JSON(http.StatusOK, user)
}

func listUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	log.Println("Fetched list of users")
	c.JSON(http.StatusOK, users)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	db.Delete(&User{}, id)
	log.Printf("User deleted: %s\n", id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/users", createUser)
	r.GET("/users", listUsers)
	r.DELETE("/users/:id", deleteUser)

	return r
}

func main() {
	initializeDB()
	r := setupRouter()
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}