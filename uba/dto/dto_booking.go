package dto

type BookingDto struct {
	Id         int     `json:"id"`
	StartDate  string  `json:"startDate"`
	EndDate    string  `json:"endDate"`
	IdMongo    string  `json:"idMongo"`
	IdUser     int     `json:"idUser"`
	TotalPrice float64 `json:"totalPrice"`
}

type BookingsDto []BookingDto
