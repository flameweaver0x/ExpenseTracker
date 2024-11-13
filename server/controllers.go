package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "strings"

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

func CreateTransactions(transactions []Transaction) error {
    valueStrings := make([]string, 0, len(transactions))
    valueArgs := make([]interface{}, 0, len(transactions)*3) // Each transaction has 3 data points
    for i, t := range transactions {
        valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
        valueArgs = append(valueArgs, t.UserID, t.Amount, t.Date)
    }
    
    stmt := fmt.Sprintf("INSERT INTO transactions(user_id, amount, date) VALUES %s", strings.Join(valueStrings, ","))
    _, err := db.Exec(stmt, valueArgs...)
    return err
}

func GetTransactions(ids []int) ([]Transaction, error) {
    valueStrings := make([]string, 0, len(ids))
    valueArgs := make([]interface{}, 0, len(ids))
    for i, id := range ids {
        valueStrings = append(valueStrings, fmt.Sprintf("$%d", i+1))
        valueArgs = append(valueArgs, id)
    }
    
    query := fmt.Sprintf("SELECT id, user_id, amount, date FROM transactions WHERE id IN (%s)", strings.Join(valueStrings, ","))
    rows, err := db.Query(query, valueArgs...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var transactions []Transaction
    for rows.Next() {
        var t Transaction
        if err := rows.Scan(&t.ID, &t.UserID, &t.Amount, &t.Date); err != nil {
            return nil, err
        }
        transactions = append(transactions, t)
    }
    
    if err = rows.Err(); err != nil {
        return nil, err
    }

    return transactions, nil
}

func main() {
    transactions := []Transaction{
        {UserID: 1, Amount: 100.50, Date: "2023-10-04"},
        {UserID: 2, Amount: 200.75, Date: "2023-10-05"},
    }
    _ = CreateTransactions(transactions)

    ids := []int{1, 2}
    transactionsRetrieved, _ := GetTransactions(ids)
    for _, t := range transactionsRetrieved {
        fmt.Printf("Transaction Retrieved: %+v\n", t)
    }
}