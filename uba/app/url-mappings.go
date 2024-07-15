package app

import (
	bookingController "uba/controllers/booking"
	hotelController "uba/controllers/hotel"
	userController "uba/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Users Mapping (listo)
	router.GET("/user/:id", userController.GetUserById)
	router.POST("/register", userController.InsertUser)
	router.POST("/login", userController.UserLogin)
	router.GET("/login", userController.UserLogin)
	router.GET("/user/email/:email", userController.GetUserByEmail)

	//Hotel Mapping (listo)
	router.POST("/hotel", hotelController.InsertHotel)
	router.GET("/hotel/:id", hotelController.GetHotelById)
	router.DELETE("/hotel/:id", hotelController.DeleteHotel)
	router.PUT("/hotel/:id", hotelController.UpdateHotel)

	//Reserva Mapping (listo)
	router.POST("/booking", bookingController.InsertBooking)
	router.GET("/available", bookingController.CheckAvailability)
	router.GET("/booking/:city", bookingController.GetAllHotelsByCity)

	log.Info("Listo el mapeo de configuraciones :)")
}
