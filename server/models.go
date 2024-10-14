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
	dsn := buildDSN()
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("Database connection successfully established")

	runMigrations()
}

func buildDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
}

func runMigrations() {
	db.AutoMigrate(&User{}, &Transaction{})
	log.Println("Database migration completed")
}

func createUser(c *gin.Context) {
	var user User
	if err := bindJSONToUser(c, &user); err != nil {
		return
	}

	createUserInDB(&user)
	c.JSON(http.StatusOK, user)
}

func bindJSONToUser(c *gin.Context, user *User) error {
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	return nil
}

func createUserInDB(user *User) {
	db.Create(user)
	log.Printf("User created: %s\n", user.Name)
}

func listUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	log.Println("Fetched list of users")
	c.JSON(http.StatusOK, users)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	deleteUserFromDB(id)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func deleteUserFromDB(id string) {
	db.Delete(&User{}, id)
	log.Printf("User deleted: %s\n", id)
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