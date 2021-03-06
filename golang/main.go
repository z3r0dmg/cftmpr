package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Stumblef00l/cftmpr/entry"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// Loading environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initializing router
	router := mux.NewRouter()

	// Add API routes here

	// User entry routes
	router.HandleFunc("/api/register", entry.RegisterUser).Methods("POST") // registers new user
	router.HandleFunc("/api/login", entry.LoginUser).Methods("POST")       // logins a user

	// Start server
	fmt.Println("Starting at port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
