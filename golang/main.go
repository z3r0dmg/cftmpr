package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	fmt.Println("Starting at port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
