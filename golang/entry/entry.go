// Package entry provides methods for endpoints for registration and login
package entry

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Stumblef00l/cftmpr/idgen"
	"github.com/Stumblef00l/cftmpr/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/pbkdf2"
)

// RegisterUser registers a new user
func RegisterUser(writer http.ResponseWriter, request *http.Request, client *mongo.Client) {

	// Get the user object from the Request

	var newUser structs.User
	_ = json.NewDecoder(request.Body).Decode(&newUser)

	// Check if the user already exists in the database
	if isRegistered(newUser.Uname, client, &newUser) {

		writer.WriteHeader(http.StatusFound)
		return
	}

	// Assign new UID to newUser
	newUser.UID = idgen.GetNewUID()

	// Hash the password
	bytePass := []byte(newUser.Pass)
	byteSalt := []byte("cftmpr")
	newUser.Pass = hashPassword(bytePass, byteSalt)

	// Assign time of joining of user
	newUser.DateJoined = time.Now().UTC()

	// Get the handle for user table
	userTable := client.Database("cftmpr").Collection("Users")

	fmt.Println(newUser)

	// Insert the newUser to the userTable
	_, err := userTable.InsertOne(context.TODO(), newUser)
	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		log.Fatal(err)
	}

	// Send confirmation of creation
	writer.WriteHeader(http.StatusOK)

}

// LoginUser implements user login
func LoginUser(writer http.ResponseWriter, request *http.Request, client *mongo.Client) {

	// Get the user object from the request
	var receivedUser structs.User
	_ = json.NewDecoder(request.Body).Decode(&receivedUser)
	var dbUser structs.User

	// Check for username registration
	if !isRegistered(receivedUser.Uname, client, &dbUser) {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	// Check for password match
	bytePass := []byte(receivedUser.Pass)
	byteSalt := []byte("cftmpr")
	receivedUser.Pass = hashPassword(bytePass, byteSalt)
	if receivedUser.Pass == dbUser.Pass {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(http.StatusUnauthorized)
	}
}

// isRegistered Returns whether a user has been registered or not
func isRegistered(uname string, client *mongo.Client, checkUser *structs.User) bool {
	// Get the handle for user table
	userTable := client.Database("cftmpr").Collection("Users")

	// Check if username is present in database
	var tempUser structs.User
	filter := bson.M{"uname": uname}
	err := userTable.FindOne(context.TODO(), filter).Decode(&tempUser)

	// Return false if no results
	if err == mongo.ErrNoDocuments {
		return false
	} else if err != nil {
		log.Fatal(err)
	}

	// Store found value in checkUser
	*checkUser = tempUser
	return true
}

// hashPassword Returns pbkdf2 hashed password
func hashPassword(password, salt []byte) string {
	return string(pbkdf2.Key(password, salt, 4096, sha256.Size, sha256.New))
}
