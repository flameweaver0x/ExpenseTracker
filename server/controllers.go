package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "strings"

    _ "github.com/lib/pq"
)

var db *sql.DB
var cache = make(map[string][]Transaction)

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
    ID       int
    UserID   int
    Amount   float64
    Date     string
    Category string
}

func CreateUser(name, email string) error {
    _, err := db.Exec("INSERT INTO users(name, email) VALUES($1, $2)", name, email)
    return err
}

func CreateTransactions(transactions []Transaction) error {
    stmt, valueArgs := prepareInsertTransactionStmt(transactions)
    _, err := db.Exec(stmt, valueArgs...)
    return err
}

func prepareInsertTransactionStmt(transactions []Transaction) (string, []interface{}) {
    valueStrings := make([]string, 0, len(transactions))
    valueArgs := make([]interface{}, 0, len(transactions)*4)
    for i, t := range transactions {
        valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4))
        valueArgs = append(valueArgs, t.UserID, t.Amount, t.Date, t.Category)
    }

    stmt := fmt.Sprintf("INSERT INTO transactions(user_id, amount, date, category) VALUES %s", strings.Join(valueStrings, ","))
    return stmt, valueArgs
}

func getCacheKey(ids []int) string {
    bytes, _ := json.Marshal(ids)
    return string(bytes)
}

func GetTransactions(ids []int) ([]Transaction, error) {
    cacheKey := getCacheKey(ids)
    if cachedTransactions, found := cache[cacheKey]; found {
        return cachedTransactions, nil
    }

    transactions, err := fetchTransactionsFromDB(ids)
    if err != nil {
        return nil, err
    }

    cache[cacheKey] = transactions
    return transactions, nil
}

func fetchTransactionsFromDB(ids []int) ([]Transaction, error) {
    query, valueArgs := prepareSelectTransactionStmt(ids)
    rows, err := db.Query(query, valueArgs...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return scanTransactions(rows)
}

func prepareSelectTransactionStmt(ids []int) (string, []interface{}) {
    valueStrings := make([]string, 0, len(ids))
    valueArgs := make([]interface{}, 0, len(ids))
    for i, id := range ids {
        valueStrings = append(valueStrings, fmt.Sprintf("$%d", i+1))
        valueArgs = append(valueArgs, id)
    }

    query := fmt.Sprintf("SELECT id, user_id, amount, date, category FROM transactions WHERE id IN (%s)", strings.Join(valueStrings, ","))
    return query, valueArgs
}

func scanTransactions(rows *sql.Rows) ([]Transaction, error) {
    var transactions []Transaction
    for rows.Next() {
        var t Transaction
        if err := rows.Scan(&t.ID, &t.UserID, &t.Amount, &t.Date, &t.Category); err != nil {
            return nil, err
        }
        transactions = append(transactions, t)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return transactions, nil
}

func UpdateTransactionCategory(id int, category string) error {
    _, err := db.Exec("UPDATE transactions SET category=$1 WHERE id=$2", category, id)
    return err
}

func main() {
    transactions := []Transaction{
        {UserID: 1, Amount: 100.50, Date: "2023-10-04", Category: "Utilities"},
        {UserID: 2, Amount: 200.75, Date: "2023-10-05", Category: "Groceries"},
    }
    err := CreateTransactions(transactions)
    if err != nil {
        log.Fatalf("Could not create transactions: %v", err)
    }

    err = UpdateTransactionCategory(1, "Entertainment")
    if err != nil {
        log.Fatalf("Could not update transaction category: %v", err)
    }

    ids := []int{1, 2}
    transactions1, err := GetTransactions(ids)
    if err != nil {
        log.Fatalf("Error retrieving transactions: %v", err)
    }
    for _, t := range transactions1 {
        fmt.Printf("First Retrieval: %+v\n", t)
    }

    transactions2, err := GetTransactions(ids)
    if err != nil {
        log.Fatalf("Error retrieving cached transactions: %v", err)
    }
    for _, t := range transactions2 {
        fmt.Printf("Second Retrieval (Cached): %+v\n", t)
    }
}