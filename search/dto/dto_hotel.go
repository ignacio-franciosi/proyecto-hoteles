package dto

type HotelDto struct {
	Id          int     `json:"id"`
	Tittle      string  `json:"tittle"`
	Description string  `json:"description"`
	Rooms       int     `json:"rooms"`
	Assessment  int     `json:"assessment"`
	Price       float32 `json:"price"`
	Gym         bool    `json:"gym"`
	Wifi        bool    `json:"wifi"`
	Parking     bool    `gorm:"parking"`
	Bidet       bool    `gorm:"bidet"`
	Pool        bool    `gorm:"pool"`
}

type HotelsDto []HotelDto

type HotelArrayDto struct {
	Id          int       `json:"id"`
	Tittle      []string  `json:"tittle"`
	Description []string  `json:"description"`
	Rooms       []int     `json:"rooms"`
	Assessment  []int     `json:"assessment"`
	Price       []float32 `json:"price"`
	Gym         []bool    `json:"gym"`
	Wifi        []bool    `json:"wifi"`
	Parking     []bool    `gorm:"parking"`
	Bidet       []bool    `gorm:"bidet"`
	Pool        []bool    `gorm:"pool"`
}

type HotelsArrayDto []HotelArrayDto
