package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

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
	// Implementation
}

func handleTransactionByID(w http.ResponseWriter, r *http.Request) {
	// Implementation
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	// Implementation
}

func handleUserByID(w http.ResponseWriter, r *http.Request) {
	// Implementation
}

func startBackgroundTasks() {
	// Background task implementation
}