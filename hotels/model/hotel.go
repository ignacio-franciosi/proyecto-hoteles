package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hotel struct {
	HotelId     primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Amenities   string             `bson:"amenities"`
	Stars       int                `bson:"stars"`
	Rooms       int                `bson:"rooms"`
	Price       float32            `bson:"price"`
	City        string             `bson:"city"`
}

type Hotels []Hotel
