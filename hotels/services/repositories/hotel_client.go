package repositories

import (
	"context"
	"fmt"
	"hotels/dto"
	"hotels/model"
	e "hotels/utils/errors"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client interface {
	GetHotelById(id string) (dto.HotelDto, e.ApiError)
	GetAllHotels() ([]dto.HotelDto, e.ApiError)
	InsertHotel(hotel dto.HotelDto) (dto.HotelDto, e.ApiError)
	UpdateHotelById(hotel dto.HotelDto) (dto.HotelDto, e.ApiError)
	DeleteHotelById(id string) e.ApiError
}

// cliente para realizar las operaciones CRUD
type HotelClient struct {
	Client     *mongo.Client   //cliente de MongoDB que se conecta a la base de datos
	Database   *mongo.Database //base de datos específica con la que interactuará el cliente
	Collection string          //colección dentro de la base de datos donde se almacenarán los datos de hoteles
}

// crea un cliente para interactuar con una base de datos MongoDB
func NewHotelnterface(host string, port int, collection string) *HotelClient {

	//crea una conexión con la base de datos MongoDB
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(fmt.Sprintf("mongodb://root:root@%s:%d/?authSource=admin&authMechanism=SCRAM-SHA-256", host, port)))
	//para conectar a MongoDB. Incluye la dirección del servidor, el puerto y la autenticación

	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	//lista los nombres de las bases de datos disponibles
	names, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	fmt.Println("[MongoDB] Initialized connection")
	fmt.Println(fmt.Sprintf("[MongoDB] Available databases: %s", names))

	return &HotelClient{
		Client:     client,
		Database:   client.Database("publicaciones"),
		Collection: collection,
	}
}

func (s *HotelClient) GetHotelById(id string) (dto.HotelDto, e.ApiError) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.HotelDto{}, e.NewBadRequestApiError(fmt.Sprintf("error getting hotel %s invalid id", id))
	}

	//realiza una consulta a la base de datos MongoDB y se busca un documento que tenga el campo "_id"
	result := s.Database.Collection(s.Collection).FindOne(context.TODO(), bson.M{"_id": objectID})

	if result.Err() == mongo.ErrNoDocuments {
		return dto.HotelDto{}, e.NewNotFoundApiError(fmt.Sprintf("hotel %s not found", id))
	}

	//si consulta exitosa, se procede a decodificar el documento en una estructura model.Hotel
	var hotel model.Hotel
	if err := result.Decode(&hotel); err != nil {
		return dto.HotelDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting hotel %s", id), err)
	}

	//devuelve dto con datos del hotel de la db
	return dto.HotelDto{
		HotelId:     id,
		Name:        hotel.Name,
		Description: hotel.Description,
		Amenities:   hotel.Amenities,
		Stars:       hotel.Stars,
		Rooms:       hotel.Rooms,
		Price:       hotel.Price,
		City:        hotel.City,
		Photos:      hotel.Photos,
	}, nil

}

// obtiene todos los hoteles de la base de datos MongoDB y los convierte a dto
// Luego, se devuelven todos los hoteles en una lista como resultado
func (s *HotelClient) GetAllHotels() (dto.HotelsDto, e.ApiError) {

	var hotelsDto dto.HotelsDto

	// Realiza una consulta para obtener todos los hoteles disponibles en la base de datos MongoDB.
	cursor, err := s.Database.Collection(s.Collection).Find(context.TODO(), bson.M{})
	if err != nil {
		return hotelsDto, e.NewInternalServerApiError("error getting hotels", err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {

		var hotel model.Hotel

		if err := cursor.Decode(&hotel); err != nil {
			return hotelsDto, e.NewInternalServerApiError("error decoding hotel", err)
		}

		// Convierte el hotel recuperado en un DTO y agrégalo a la lista de hoteles.
		hotelDto := dto.HotelDto{
			HotelId:     hotel.ID.Hex(),
			Name:        hotel.Name,
			Description: hotel.Description,
			Amenities:   hotel.Amenities,
			Stars:       hotel.Stars,
			Rooms:       hotel.Rooms,
			Price:       hotel.Price,
			City:        hotel.City,
			Photos:      hotel.Photos,
		}

		hotelsDto = append(hotelsDto, hotelDto)
	}

	if err := cursor.Err(); err != nil {
		return hotelsDto, e.NewInternalServerApiError("error iterating over hotels", err)
	}

	return hotelsDto, nil
}

func (s *HotelClient) InsertHotel(hotel dto.HotelDto) (dto.HotelDto, e.ApiError) {

	id := primitive.NewObjectID()

	//Se crea un nuevo documento con los campos especificados y se inserta en la base de datos MongoDB
	result, err := s.Database.Collection(s.Collection).InsertOne(context.TODO(), model.Hotel{
		HotelId:     id,
		Name:        hotel.Name,
		Description: hotel.Description,
		Amenities:   hotel.Amenities,
		Stars:       hotel.Stars,
		Rooms:       hotel.Rooms,
		Price:       hotel.Price,
		City:        hotel.City,
		Photos:      hotel.Photos,
	})

	if err != nil {
		return hotel, e.NewInternalServerApiError(fmt.Sprintf("error inserting to mongo %s", hotel.HotelId), err)
	}
	hotel.HotelId = fmt.Sprintf(result.InsertedID.(primitive.ObjectID).Hex())

	return hotel, nil
}

func (s *HotelClient) DeleteHotel(id string) e.ApiError {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return e.NewBadRequestApiError(fmt.Sprintf("error deleting hotel %s invalid id", id))
	}

	result, err := s.Database.Collection(s.Collection).DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		log.Error(err)
		return e.NewInternalServerApiError("error deleting hotel", err)
	}
	log.Debug(result.DeletedCount)

	result, err = s.Database.Collection(s.Collection).DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		log.Error(err)
		return e.NewInternalServerApiError("error deleting hotel", err)
	}
	log.Debug(result.DeletedCount)

	return nil
}

func (s *HotelClient) UpdateHotelById(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {

	//convierte el ID del hotel en un objeto ObjectID.
	objectID, err := primitive.ObjectIDFromHex(hotelDto.HotelId)
	if err != nil {
		return dto.HotelDto{}, e.NewBadRequestApiError(fmt.Sprintf("error updating hotel %s invalid id", hotelDto.HotelId))
	}

	//actualiza el hotel en la base de datos.
	update := bson.M{
		"$set": bson.M{
			"Name":        hotelDto.Name,
			"Description": hotelDto.Description,
			"Amenities":   hotelDto.Amenities,
			"Stars":       hotelDto.Stars,
			"Rooms":       hotelDto.Rooms,
			"Price":       hotelDto.Price,
			"City":        hotelDto.City,
			"Photos":      hotelDto.Photos,
		},
	}

	//Realiza la actualización en la base de datos.
	filter := bson.M{"_id": objectID}
	_, err = s.Database.Collection(s.Collection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return dto.HotelDto{}, e.NewInternalServerApiError("error updating hotel", err)
	}

	//Verifica si la actualización fue exitosa.
	if err != nil {
		return dto.HotelDto{}, e.NewBadRequestApiError("error in update")
	}

	return hotelDto, nil
}
