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
	router.GET("/user/email", userController.GetUserByEmail)

	//Hotel Mapping (listo)
	router.POST("/hotel", hotelController.InsertHotel)
	router.GET("/hotel/:id", hotelController.GetHotelById)

	//Reserva Mapping (listo)
	router.POST("/booking", bookingController.InsertBooking)

	log.Info("Listo el mapeo de configuraciones :)")
}
