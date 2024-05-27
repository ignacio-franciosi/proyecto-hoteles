package booking

import (
	"strings"
	"uba/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InsertBooking(booking model.Booking) model.Booking {
	result := Db.Create(&booking)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Reserva Creada: ", booking.Id)
	return booking
}

func GetBookingsByHotel(idMongo string) model.Bookings {
	var bookings model.Bookings

	Db.Where("id_mongo = ?", idMongo).Find(&bookings)
	log.Debug("Bookings: ", bookings)

	return bookings
}

func GetAllHotelsByCity(city string) model.Hotels {

	cityFormatted := strings.ReplaceAll(city, "+", " ")
	var hotels model.Hotels

	Db.Where("city = ?", cityFormatted).Find(&hotels)

	return hotels
}
