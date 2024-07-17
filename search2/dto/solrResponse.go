package dto

type SolrResponseDto struct {
	Id          string    `json:"id"`
	Name        []string  `json:"name" validate:"required"`
	Rooms       []int     `json:"rooms" validate:"required"`
	Description []string  `json:"description" validate:"required"`
	City        []string  `json:"city" validate:"required"`
	Stars       []int     `json:"stars" validate:"required"`
	Price       []float64 `json:"price" validate:"required"`
	Amenities   []string  `json:"amenities" validate:"required"`
}

type SolrResponsesDto []SolrResponseDto
