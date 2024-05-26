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

	client.Disconnect(context.TODO())
}

func Init_db() {

	//crea una conexión con la base de datos MongoDB
	clientOpts := options.Client().ApplyURI("mongodb://root:pass@mongodatabase:27017/?authSource=admin&authMechanism=SCRAM-SHA-256")
	cli, err := mongo.Connect(context.TODO(), clientOpts)
	client = cli

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	//lista los nombres de las bases de datos disponibles
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

/*

type HotelClient struct {
	Client     *mongo.Client   //cliente de MongoDB que se conecta a la base de datos
	Database   *mongo.Database //base de datos específica con la que interactuará el cliente
	Collection string          //colección dentro de la base de datos donde se almacenarán los datos de hoteles
}

// crea un cliente para interactuar con una base de datos MongoDB
func NewHotelInterface(host string, port int, collection string) *HotelClient {

	//crea una conexión con la base de datos MongoDB
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://root:pass@mongodatabase:27017/?authSource=admin&authMechanism=SCRAM-SHA-256"))

	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	//lista los nombres de las bases de datos disponibles
	names, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	fmt.Println("[MongoDB] Initialized connection")
	fmt.Println(names)

	return &HotelClient{
		Client:     client,
		Database:   client.Database("hotels_db"),
		Collection: collection,
	}
}


*/
