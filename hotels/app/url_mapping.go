package app

import (
	hotelController "hotels/controllers"
	//"Hotel/controller"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Hotels Mapping

	router.GET("/hotels/:HotelId", hotelController.GetHotelById)

	router.GET("/hotels", hotelController.GetAllHotels)

	router.POST("/hotel", hotelController.InsertHotel)

	//router.POST("/hotels", hotelController.QueueHotels)

	router.PUT("/hotel/:HotelId", hotelController.UpdateHotel)

	//router.PUT("/hotels", hotelController.QueueHotels)

	router.DELETE("/hotel/:HotelId", hotelController.DeleteHotel)

	log.Info("Finishing mappings configurations")
}
