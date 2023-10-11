package app

import (
	hotelController "hotels/controllers/hotel"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Hotels Mapping
	
	router.GET("/hotels/:HotelId", hotelController.GetHotelById)
	
	router.POST("/hotel", hotelController.InsertHotel)
	// router.POST("/hotels", hotelController.QueueHotels)

	router.PUT("/hotel/:HotelId", hotelController.UpdateHotelById)
	// router.PUT("/hotels", hotelController.QueueHotels)

	router.DELETE("/hotel/:HotelId", hotelController.DeleteHotelById)

	log.Info("Finishing mappings configurations")
}