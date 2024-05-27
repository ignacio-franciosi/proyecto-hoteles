package services

import (
	hotelClient "uba/clients/hotel"
	"uba/dto"
	"uba/model"
	e "uba/utils/errors"
)

type hotelService struct{}

type hotelServiceInterface interface {
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
	GetHotelById(id string) (dto.HotelDto, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {
	var hotel model.Hotel

	hotel.Id = hotelDto.Id
	hotel.IdAmadeus = hotelDto.IdAmadeus
	hotel.IdMongo = hotelDto.IdMongo
	hotel.Rooms = hotelDto.Rooms
	hotel.Price = hotelDto.Price
	hotel.City = hotelDto.City

	hotel = hotelClient.InsertHotel(hotel)

	var response dto.HotelDto

	response.Id = hotel.Id
	response.IdAmadeus = hotel.IdAmadeus
	response.IdMongo = hotel.IdMongo
	response.Rooms = hotel.Rooms
	response.Price = hotel.Price
	response.City = hotel.City

	return response, nil
}

func (s *hotelService) GetHotelById(id string) (dto.HotelDto, e.ApiError) {
	var hotel model.Hotel = hotelClient.GetHotelById(id)
	var hotelDto dto.HotelDto

	if hotel.IdMongo == "000000000000000000000000" {
		return hotelDto, e.NewBadRequestApiError("Hotel no encontrado")
	}

	hotelDto.Id = hotel.Id
	hotelDto.IdMongo = hotel.IdMongo
	hotelDto.IdAmadeus = hotel.IdAmadeus
	hotelDto.Rooms = hotel.Rooms
	hotelDto.Price = hotel.Price
	hotelDto.City = hotel.City
	return hotelDto, nil

}
