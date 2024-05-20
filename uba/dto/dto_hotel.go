package dto

type HotelDto struct {
	Id        int     `json:"id"`
	IdMongo   string  `json:"id_mongo"`
	IdAmadeus string  `json:"id_amadeus"`
	Rooms     int     `json:"rooms"`
	Price     float64 `json:"price"`
	City      string  `json:"city"`
}

type HotelsDto []HotelDto
