package dto

import "time"

type BookingDto struct {
	Id         int       `json:"id"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	TotalPrice float32   `json:"totalPrice"`
	IdHotel    int       `json:"idHotel"`
	IdUser     int       `json:"idUser"`
}

type BookingsDto []BookingDto
