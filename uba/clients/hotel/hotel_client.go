package hotel

import (
	"uba/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InsertHotel(hotel model.Hotel) model.Hotel {
	result := Db.Create(&hotel)

	if result.Error != nil {
		log.Error("")
	}

	log.Debug("Hotel Created: ", hotel.Id)
	return hotel
}

func GetHotelById(id string) model.Hotel {
	var hotel model.Hotel

	Db.Where("id_mongo = ?", id).First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}
