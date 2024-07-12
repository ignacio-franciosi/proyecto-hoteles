package app

import (
	"search2/controller"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Add all methods and its mappings
	router.GET("/hotel/:id", controller.GetHotelById)
	router.GET("/hotel", controller.GetHotels) //se le puede agregar city a este url, ej: http://localhost:8000/hotels?city=Paris

	log.Info("Finishing mappings configurations")
}
