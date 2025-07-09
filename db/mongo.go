package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MessageColl, SentColl *mongo.Collection
var mongoClient *mongo.Client

func InitMongo() (err error) {
	// Start MongoDB Connection
	mongoClient, err = mongo.Connect(options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@mongo:27017", os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD"))))

	if err != nil {
		return
	}

	db := mongoClient.Database("main")

	// Check if required collections exists
	if err = checkColls(db); err != nil {
		return
	}

	// Set exported collection
	MessageColl = db.Collection("messages")

	return
}

func checkColls(db *mongo.Database) (err error) {
	// Get collections from db
	colls, err := db.ListCollectionNames(context.TODO(), bson.M{})

	if err != nil {
		return
	}

	// Check if messages collection exists
	exists := false
	for _, v := range colls {
		if v == "messages" {
			exists = true
			break
		}
	}

	// Create messages collection if it doesn't exists
	if !exists {
		if err = db.CreateCollection(context.TODO(), "messages"); err != nil {
			return
		}

		// Create index for the new collection
		if _, err = db.Collection("messages").Indexes().CreateOne(context.TODO(), mongo.IndexModel{Keys: map[string]any{"sent": 1}}); err != nil {
			return
		}
	}

	return
}

func DisconnectMongo() {
	mongoClient.Disconnect(context.TODO())
}
