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

	//crea una conexión con la base de datos MongoDB
	//clientOpts := options.Client().ApplyURI("mongodb://root:pass@mongodatabase:27017/?authSource=admin&authMechanism=SCRAM-SHA-256")
	//clientOpts := options.Client().ApplyURI("mongodb://mongo:27017")
	//clientOpts := options.Client().ApplyURI("mongodb://root:pass@mongo:27017/?authSource=admin&authMechanism=SCRAM-SHA-256")
	clientOpts := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.10")
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

//clientOpts := options.Client().ApplyURI("mongodb://admin:pass@mongodatabase:27017/?authSource=admin&authMechanism=SCRAM-SHA-256")
/*clientOpts := options.Client().ApplyURI("mongodb://mongo:27017").
SetAuth(options.Credential{
	AuthSource:    "admin",         // Reemplaza "admin" con el nombre de la base de datos de autenticación que desees utilizar.
	AuthMechanism: "SCRAM-SHA-256", // Reemplaza con el mecanismo de autenticación adecuado si no es el predeterminado.
	Username:      "admin",         // Reemplaza con tu nombre de usuario.
	Password:      "pass",          // Reemplaza con tu contraseña.
})*/

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
