package dto

type HotelDto struct {
	HotelId     string  `json:"hotel_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Amenities   string  `json:"amenities"`
	Stars       int     `json:"stars"`
	Rooms       int     `json:"rooms"`
	Price       float32 `json:"price"`
	City        string  `json:"city"`
}
type HotelsDto []HotelDto
