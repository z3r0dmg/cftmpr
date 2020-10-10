// Package idgen creates and fetches new random IDs
package idgen

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Stumblef00l/cftmpr/structs"
)

// GetNewUID gets a new user ID
func GetNewUID() string {

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

	// Get UserIdStore collection handle
	userIDStore := client.Database("cftmpr").Collection("UserIdStore")

	// Get One new user ID
	filter := bson.M{"used": false}
	var newUser structs.UserIDStruct
	err = userIDStore.FindOne(context.TODO(), filter).Decode(&newUser)

	// Check for errors
	if err != nil {
		log.Fatal(err)
	}

	// Return user ID
	return newUser.UID
}

// GetNewSessID gets a new user ID
func GetNewSessID() string {

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

	// Get SessIdStore collection handle
	sessIDStore := client.Database("cftmpr").Collection("SessIdStore")

	// Get One new sessID
	filter := bson.M{"used": false}
	var newSess structs.SessIDStruct
	err = sessIDStore.FindOne(context.TODO(), filter).Decode(&newSess)

	// Check for errors
	if err != nil {
		log.Fatal(err)
	}

	// Return user ID
	return newSess.SessID
}
