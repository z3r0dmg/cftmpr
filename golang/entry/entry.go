// Package entry provides methods for endpoints for registration and login
package entry

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Stumblef00l/cftmpr/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RegisterUser registers a new user
func RegisterUser(writer http.ResponseWriter, request *http.Request) {
	// Secret URI
	uri := os.Getenv("CFTMPR_ATLAS_URI")

	// Connecting to database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect after query
	defer client.Disconnect(ctx)

	// Get the user object from the Request

	var newUser structs.User
	_ = json.NewDecoder(request.Body).Decode(&newUser)

	// Set response type
	writer.Header().Set("Content-Type", "application/json")

	// Check if the user already exists in the database
	if isRegistered(newUser.Uname, client) {

		writer.WriteHeader(http.StatusFound)
		return
	}

	// Get the handle for user table
	userTable := client.Database("cftmpr").Collection("Users")

	// Insert the newUser to the userTable
	result, err := userTable.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(writer).Encode(result)

}

// LoginUser verifies user credentials upon login request
func LoginUser(writer http.ResponseWriter, request *http.Request) {

}

// isRegistered Returns whether a user has been registered or not
func isRegistered(uname string, client *mongo.Client) bool {
	// Get the handle for user table
	userTable := client.Database("cftmpr").Collection("Users")

	// Get One new user ID
	filter := bson.M{"uname": uname}
	var user structs.User
	err := userTable.FindOne(context.TODO(), filter).Decode(&user)

	if err == mongo.ErrNoDocuments {
		return false
	} else if err != nil {
		log.Fatal(err)
	}

	return true
}
