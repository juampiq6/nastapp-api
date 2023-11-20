package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

var dbName string
var mongoConnectionString string

func InitMongoClient() {
	getEnvConfig()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoConnectionString).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.TODO(), opts.ReadPreference)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("[Mongo-Connection] Ping to Mongo DB successful!")
	}
	mongoClient = client
}

func getEnvConfig() {
	if dbName == "" {
		dbName = os.Getenv("DB_NAME")
		// if non set on .env, then default is "nastapp"
		if dbName == "" {
			dbName = "nastapp"
		}
	}
	mongoConnectionString = os.Getenv("MONGO_STRING")
	if mongoConnectionString == "" {
		panic("You must set your 'MONGODB_URI' environment variable.")
	}
}

func getMongoClientInstance() *mongo.Client {
	return mongoClient
}

func KillMongoClient() {
	var err = mongoClient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
}
