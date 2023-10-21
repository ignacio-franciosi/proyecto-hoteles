package hotel

import (
	"uba/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
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
