package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "log"
    "net/http"
    "os"
    "sync"
)

type Transaction struct {
    ID          string  `json:"id"`
    Description string  `json:"description"`
    Amount      float64 `json:"amount"`
}

type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

var transactions = struct {
    sync.RWMutex
    items []Transaction
}{}

var users = struct {
    sync.RWMutex
    items []User
}{}

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    router := setupRouter()
    port := getServerPort()

    go startBackgroundTasks()

    log.Printf("Server starting on port %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}

func setupRouter() *mux.Router {
    router := mux.NewRouter()
    setupRoutes(router.PathPrefix("/api").Subrouter())
    return router
}

func setupRoutes(api *mux.Router) {
    api.HandleFunc("/transactions", handleTransactions).Methods(http.MethodGet, http.MethodPost)
    api.HandleFunc("/transactions/{id}", handleTransactionByID).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
    api.HandleFunc("/users", handleUsers).Methods(http.MethodGet, http.MethodPost)
    api.HandleFunc("/users/{id}", handleUserByID).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
}

func getServerPort() string {
    port := os.Getenv("SERVER_PORT")
    if port == "" {
        port = "8080"
    }
    return port
}

func handleTransactions(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        transactions.RLock()
        json.NewEncoder(w).Encode(transactions.items) // Handle errors in real code
        transactions.RUnlock()
    case http.MethodPost:
        var transaction Transaction
        if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        transactions.Lock()
        transactions.items = append(transactions.items, transaction)
        transactions.Unlock()
        w.WriteHeader(http.StatusCreated)
    }
}

func handleTransactionByID(w http.ResponseWriter, r *http.Request) {
    // Implementation
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        users.RLock()
        json.NewEncoder(w).Encode(users.items) // Handle errors in real code
        users.RUnlock()
    case http.MethodPost:
        var user User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        users.Lock()
        users.items = append(users.items, user)
        users.Unlock()
        w.WriteHeader(http.StatusCreated)
    }
}

func handleUserByID(w http.ResponseWriter, r *http.Request) {
    // Implementation
}

func startBackgroundTasks() {
    // Background task implementation
}