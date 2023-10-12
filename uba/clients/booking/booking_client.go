package booking

//ORM traductor
import (
	"repo/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InsertReserva(reserva model.Booking) model.Booking {
	result := Db.Create(&booking)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Reserva creada: ", booking.Id)
	return booking
}
