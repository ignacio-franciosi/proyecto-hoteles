package app

import (
	hotelController "hotels/controllers"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Hotels Mapping

	router.GET("/hotels/:HotelId", hotelController.GetHotelById)

	router.GET("/hotels", hotelController.GetAllHotels)

	router.POST("/hotel", hotelController.InsertHotel)

	router.PUT("/hotel/:HotelId", hotelController.UpdateHotel)

	router.DELETE("/hotel/:HotelId", hotelController.DeleteHotel)

	log.Info("Finishing mappings configurations")
}
