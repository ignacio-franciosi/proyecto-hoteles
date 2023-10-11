package dto

type HotelDto struct {
	
	HotelId      	string 	`json:"hotel_id"`
	Name      		string 	`json:"name"`
	Description     string 	`json:"description"`
	Amenities   	string 	`json:"amenities"`
	Stars 			string 	`json:"stars"`
	Rooms 			int    	`json:"rooms"`
	Price  			float32 `json:"price"`
	City      		string 	`json:"city"`
	Photos			string 	`json:"photos"`
}
type HotelsDto []HotelDto