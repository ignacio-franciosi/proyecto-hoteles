package dto

type BookingDto struct {
	Id         int     `json:"id"`
	StartDate  int     `json:"startDate"`
	EndDate    int     `json:"endDate"`
	TotalPrice float32 `json:"totalPrice"`
	IdHotel    int     `json:"idHotel"`
	IdUser     int     `json:"idUser"`
}

type BookingsDto []BookingDto
