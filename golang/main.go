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

	var client *mongo.Client
	client = nil
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// helpers will calls to the actual handlers
	RegisterUserHandler := func(w http.ResponseWriter, r *http.Request) {
		if client == nil {
			// Secret URI
			uri := os.Getenv("CFTMPR_ATLAS_URI")

			client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
			if err != nil {
				log.Fatal(err)
			}
		}
		entry.RegisterUser(w, r, client)
	}

	LoginUserHandler := func(w http.ResponseWriter, r *http.Request) {
		if client == nil {
			// Secret URI
			uri := os.Getenv("CFTMPR_ATLAS_URI")

			client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
			if err != nil {
				log.Fatal(err)
			}
		}
		entry.LoginUser(w, r, client)
	}

	// Add API routes here

	// User entry routes
	router.HandleFunc("/api/register", RegisterUserHandler).Methods("POST") // registers new user
	router.HandleFunc("/api/login", LoginUserHandler).Methods("POST")       // logins a user

	// Add API routes here

	// User entry routes
	router.HandleFunc("/api/register", entry.RegisterUser).Methods("POST") // registers new user
	router.HandleFunc("/api/login", entry.LoginUser).Methods("POST")       // logins a user

	// Start server
	fmt.Println("Starting at port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))

	client.Disconnect(ctx)
}
