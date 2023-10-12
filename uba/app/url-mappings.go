package app

import (
	hotelController "repo/controllers/hotel"
	reservaController "repo/controllers/reserva"
	userController "repo/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Users Mapping (listo)
	router.GET("/user/:id", userController.GetUserById)
	router.POST("/register", userController.InsertUser)
	router.POST("/login", userController.UserLogin)
	router.GET("/login", userController.UserLogin)
	router.GET("/user/email", userController.GetUserByEmail)

	//Hotel Mapping (ignorar)
	router.GET("/hotel/:id/habitaciones", hotelController.GetHotelDisponibilidad)
	router.GET("/hotel/:id", hotelController.GetHotelById)
	router.GET("/hotel", hotelController.GetHoteles)
	router.POST("/hotel", hotelController.InsertHotel)

	//Reserva Mapping (pendiente)
	router.POST("/reserva", reservaController.ReservaInsert)
	router.GET("/reservaUser/:idUser", reservaController.GetReservasByIdUser)

	log.Info("Listo el mapeo de configuraciones :)")
}
