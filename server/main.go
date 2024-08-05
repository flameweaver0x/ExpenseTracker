package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/transactions", handleTransactions).Methods("GET", "POST")
	router.HandleFunc("/api/transactions/{id}", handleTransactionByID).Methods("GET", "PUT", "DELETE")
	router.HandleFunc("/api/users", handleUsers).Methods("GET", "POST")
	router.HandleFunc("/api/users/{id}", handleUserByID).Methods("GET", "PUT", "DELETE")
	
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	go startBackgroundTasks()

	log.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

func handleTransactions(w http.ResponseWriter, r *http.Request) {
}

func handleTransactionByID(w http.ResponseWriter, r *http.Request) {
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
}

func handleUserByID(w http.ResponseWriter, r *http.Request) {
}

func startBackgroundTasks() {
}