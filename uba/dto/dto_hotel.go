package dto

type HotelDto struct {
	IdHotel   int     `json:"id_hotel"`
	IdMongo   string  `json:"id_mongo"`
	IdAmadeus string  `json:"id_amadeus"`
	Rooms     int     `json:"rooms"`
	Price     float64 `json:"price"`
	City      string  `json:"city"`
}

type HotelsDto []HotelDto
