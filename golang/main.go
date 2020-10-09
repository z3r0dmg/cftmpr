package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Stumblef00l/cftmpr/idgen"
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
	newID := idgen.GetNewUID()

	fmt.Println(newID)

	fmt.Println("Starting at port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
