package dto

type InsertHotelDto struct {
	Id      int     `json:"id"`
	IdMongo string  `json:"id_mongo"`
	City    string  `json:"city"`
	Price   float32 `json:"price"`
}
