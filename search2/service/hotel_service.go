package service

import (
	"encoding/json"
	"search2/client"
	"search2/dto"

	log "github.com/sirupsen/logrus"
)

type hotelService struct{}

type hotelServiceInterface interface {
	InsertUpdateHotel(hotelDto dto.HotelDto) error
	GetHotels() (dto.HotelsDto, error)
	GetHotelById(id string) (dto.HotelDto, error)
	GetHotelByCity(id string) (dto.HotelsDto, error)
	DeleteHotelById(id string) error
}

var HotelService hotelServiceInterface

func init() {
	HotelService = &hotelService{}
}

func (s hotelService) InsertUpdateHotel(hotelDto dto.HotelDto) error {

	document := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{
				"id":          hotelDto.HotelId,
				"name":        hotelDto.Name,
				"rooms":       hotelDto.Rooms,
				"description": hotelDto.Description,
				"city":        hotelDto.City,
				"stars":       hotelDto.Stars,
				"price":       hotelDto.Price,
				"amenities":   hotelDto.Amenities,
			},
		},
	}

	err := client.SolrHotelClient.UpdateHotel(document)

	if err != nil {
		log.Info("Error updating hotel", err)
		return err
	}

	return nil
}

func (s hotelService) GetHotels() (dto.HotelsDto, error) {

	var solrResponsesDto dto.SolrResponsesDto
	results, err := client.SolrHotelClient.GetHotels()

	if err != nil {
		log.Info(err)
		return dto.HotelsDto{}, err
	}

	for i := 0; i < results.Len(); i++ {
		var solrResponseDto dto.SolrResponseDto

		jsonResult, err := json.Marshal(results.Get(i).Fields)

		if err != nil {
			return dto.HotelsDto{}, err
		}

		err = json.Unmarshal(jsonResult, &solrResponseDto)

		if err != nil {
			return dto.HotelsDto{}, err
		}

		solrResponsesDto = append(solrResponsesDto, solrResponseDto)
	}

	hotelsDto := unmarshalSolrResponse(solrResponsesDto)

	return hotelsDto, nil
}

func (s hotelService) GetHotelById(id string) (dto.HotelDto, error) {

	var solrResponsesDto dto.SolrResponsesDto
	results, err := client.SolrHotelClient.GetHotelById(id)

	if err != nil {
		log.Info(err)
		return dto.HotelDto{}, err
	}

	for i := 0; i < results.Len(); i++ {
		var solrResponseDto dto.SolrResponseDto

		jsonResult, err := json.Marshal(results.Get(i).Fields)

		if err != nil {
			return dto.HotelDto{}, err
		}

		err = json.Unmarshal(jsonResult, &solrResponseDto)

		if err != nil {
			return dto.HotelDto{}, err
		}

		solrResponsesDto = append(solrResponsesDto, solrResponseDto)
	}

	hotelsDto := unmarshalSolrResponse(solrResponsesDto)

	return hotelsDto[0], nil
}

func (s hotelService) GetHotelByCity(city string) (dto.HotelsDto, error) {

	var solrResponsesDto dto.SolrResponsesDto
	results, err := client.SolrHotelClient.GetHotelsByCity(city)

	if err != nil {
		log.Info(err)
		return dto.HotelsDto{}, err
	}

	for i := 0; i < results.Len(); i++ {
		var solrResponseDto dto.SolrResponseDto

		jsonResult, err := json.Marshal(results.Get(i).Fields)

		if err != nil {
			return dto.HotelsDto{}, err
		}

		err = json.Unmarshal(jsonResult, &solrResponseDto)

		if err != nil {
			return dto.HotelsDto{}, err
		}

		solrResponsesDto = append(solrResponsesDto, solrResponseDto)
	}

	hotelsDto := unmarshalSolrResponse(solrResponsesDto)

	return hotelsDto, nil

}

func (s hotelService) DeleteHotelById(id string) error {

	document := map[string]interface{}{
		"delete": []interface{}{
			map[string]interface{}{
				"id": id,
			},
		},
	}

	err := client.SolrHotelClient.UpdateHotel(document)

	if err != nil {
		log.Info("Error deleting hotel", err)
		return err
	}

	return nil
}

func unmarshalSolrResponse(responses dto.SolrResponsesDto) dto.HotelsDto {
	var hotelsDto dto.HotelsDto

	for _, response := range responses {
		var hotelDto dto.HotelDto

		hotelDto.HotelId = response.Id
		hotelDto.Name = response.Name[0]
		hotelDto.Rooms = response.Rooms[0]
		hotelDto.Description = response.Description[0]
		hotelDto.City = response.City[0]
		hotelDto.Stars = response.Stars[0]
		hotelDto.Price = response.Price[0]
		hotelDto.Amenities = response.Amenities[0]

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto
}
