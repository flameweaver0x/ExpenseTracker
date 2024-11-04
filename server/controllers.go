package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %s", err)
	}
}

type Transaction struct {
	ID     int
	UserID int
	Amount float64
	Date   string
}

func CreateUser(name string, email string) error {
	_, err := db.Exec("INSERT INTO users(name, email) VALUES($1, $2)", name, email)
	return err
}

func CreateTransaction(t Transaction) error {
	_, err := db.Exec("INSERT INTO transactions(user_id, amount, date) VALUES($1, $2, $3)", t.UserID, t.Amount, t.Date)
	return err
}

func GetTransaction(id int) (*Transaction, error) {
	t := &Transaction{}
	err := db.QueryRow("SELECT id, user_id, amount, date FROM transactions WHERE id = $1", id).Scan(&t.ID, &t.UserID, &t.Amount, &t.Date)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func UpdateTransaction(t Transaction) error {
	_, err := db.Exec("UPDATE transactions SET user_id = $1, amount = $2, date = $3 WHERE id = $4", t.UserID, t.Amount, t.Date, t.ID)
	return err
}

func DeleteTransaction(id int) error {
	_, err := db.Exec("DELETE FROM transactions WHERE id = $1", id)
	return err
}

func main() {
	err := CreateUser("John Doe", "john@example.com")
	if err != nil {
		log.Fatalf("Error creating new user: %s", err)
	}

	t := Transaction{UserID: 1, Amount: 100.50, Date: "2023-10-04"}
	err = CreateTransaction(t)
	if err != nil {
		log.Fatalf("Error creating transaction: %s", err)
	}

	transaction, err := GetTransaction(1)
	if err != nil {
		log.Fatalf("Error getting transaction: %s", err)
	}
	fmt.Println("Transaction retrieved:", transaction)

	err = UpdateTransaction(Transaction{ID: 1, UserID: 1, Amount: 150.0, Date: "2023-10-05"})
	if err != nil {
		log.Fatalf("Error updating transaction: %s", err)
	}

	err = DeleteTransaction(1)
	if err != nil {
		log.Fatalf("Error deleting transaction: %s", err)
	}
}