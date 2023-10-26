package dto

type BookingDto struct {
	Id         int     `json:"id"`
	StartDate  string  `json:"startDate"`
	EndDate    string  `json:"endDate"`
	TotalPrice float32 `json:"totalPrice"`
	IdHotel    int     `json:"idHotel"`
	IdUser     int     `json:"idUser"`
}

type BookingsDto []BookingDto
