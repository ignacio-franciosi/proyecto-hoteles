package db

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDb *mongo.Database
var client *mongo.Client
var HotelsCollection *mongo.Collection

func Disconect_db() {

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Disconnected from MongoDB")
}

func Init_db() {

	//clientOpts := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.10")
	clientOpts := options.Client().ApplyURI("mongodb://mongo:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.10")

	cli, err := mongo.Connect(context.TODO(), clientOpts)
	client = cli

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	dbNames, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		log.Info("Failed to get databases available")
		log.Fatal(err)
	}

	MongoDb = client.Database("hotels_db")

	fmt.Println("Available datatabases:")
	fmt.Println(dbNames)

	HotelsCollection = MongoDb.Collection("hotels")

}
