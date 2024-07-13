package client

import (
	"context"
	"errors"
	"fmt"
	db "hotels/db"
	"hotels/model"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type hotelClient struct{}

type hotelClientInterface interface {
	InsertHotel(hotel model.Hotel) model.Hotel
	GetHotelById(id string) (model.Hotel, error)
	GetAllHotels() model.Hotels
	DeleteHotelById(id string) error
	UpdateHotelById(id string, hotel model.Hotel) model.Hotel
}

var HotelClient hotelClientInterface

func init() {
	HotelClient = &hotelClient{}
}

func (c hotelClient) InsertHotel(hotel model.Hotel) model.Hotel {

	insertHotel := hotel
	insertHotel.HotelId = primitive.NewObjectID()

	_, err := db.HotelsCollection.InsertOne(context.TODO(), &insertHotel)

	if err != nil {
		fmt.Println(err)
		return hotel
	}

	hotel.HotelId = insertHotel.HotelId

	return hotel
}

func (c hotelClient) GetHotelById(id string) (model.Hotel, error) {
	var hotel model.Hotel

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return hotel, err
	}

	err = db.HotelsCollection.FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&hotel)
	if err != nil {
		return hotel, err
	}
	return hotel, nil
}

func (c hotelClient) GetAllHotels() model.Hotels {
	var hotels model.Hotels

	cursor, err := db.HotelsCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		fmt.Println(err)
		return hotels
	}

	err = cursor.All(context.TODO(), &hotels)

	if err != nil {
		fmt.Println(err)
		return hotels
	}

	return hotels
}

func (c hotelClient) DeleteHotelById(id string) error {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := db.HotelsCollection.DeleteOne(context.TODO(), bson.D{{"_id", objID}})

	if err != nil {
		log.Debug("Failed to delete hotel")
		return err
	} else if result.DeletedCount == 0 {
		log.Debug("Hotel not found")
		return errors.New("hotel not found")
	}

	log.Debug("Hotel deleted successfully: ", id)
	return nil
}

func (c hotelClient) UpdateHotelById(id string, hotel model.Hotel) model.Hotel {

	db := db.MongoDb

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Debug("Failed to convert hex ID to ObjectID")
		return model.Hotel{}
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: hotel.Name},
			{Key: "description", Value: hotel.Description},
			{Key: "amenities", Value: hotel.Amenities},
			{Key: "stars", Value: hotel.Stars},
			{Key: "rooms", Value: hotel.Rooms},
			{Key: "city", Value: hotel.City},
			{Key: "price", Value: hotel.Price},
		}},
	}

	result, err := db.Collection("hotels").UpdateOne(context.TODO(), bson.D{{"_id", objID}}, update)
	if err != nil {
		log.Debug("Failed to update hotel")
		return model.Hotel{}
	}

	if result.MatchedCount != 0 {
		log.Debug("Updated hotel successfully")
		return hotel
	}

	log.Debug("No matching hotel found to update")
	return model.Hotel{}
}
