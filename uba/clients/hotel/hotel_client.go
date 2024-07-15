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
		log.Error("Error creating hotel: ", result.Error)
	}

	log.Debug("Hotel Created: ", hotel.IdHotel)
	return hotel
}

func GetHotelById(id string) model.Hotel {
	var hotel model.Hotel

	Db.Where("id_mongo = ?", id).First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

func DeleteHotel(hotel model.Hotel) error {
	err := Db.Delete(&hotel).Error

	if err != nil {
		log.Debug("Failed to delete hotel")
	} else {
		log.Debug("Hotel deleted: ", hotel.IdHotel)
	}
	return err
}

func UpdateHotel(hotel model.Hotel) model.Hotel {

	result := Db.Save(&hotel)

	if result.Error != nil {
		log.Debug("Failed to update hotel")
		return model.Hotel{}
	}

	log.Debug("Updated hotel: ", hotel.IdHotel)
	return hotel
}
