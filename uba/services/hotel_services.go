package service

import (
	hotelClient "uba/clients/hotel"
	"uba/dto"
	"uba/model"
	e "uba/utils/errors"
)

type hotelService struct{}

type hotelServiceInterface interface {
	InsertHotel(hotelDto dto.HotelPostDto, idAmadeus string) (dto.HotelDto, e.ApiError)
	GetHotelById(id int) (dto.HotelDto, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) InsertHotel(hotelDto dto.InsertHotelDto, idAmadeus string) (dto.HotelDto, e.ApiError) {
	var hotel model.Hotel
	var response dto.HotelDto

	hotel.Price = hotelDto.Price
	hotel.IdAmadeus = idAmadeus
	hotel.IdMongo = hotelDto.IdMongo

	hotel = hotelClient.InsertHotel(hotel)

	hotelDto.Id = hotel.Id

	return response, nil
}

func (s *hotelService) GetHotelById(id int) (dto.HotelDto, e.ApiError) {
	var hotel model.Hotel = hotelClient.GetHotelById(id)
	var hotelDto dto.HotelDto

	if hotel.Id == 0 {
		return hotelDto, e.NewBadRequestApiError("Hotel no encontrado")
	}

	hotelDto.Id = hotel.Id
	hotelDto.Price = hotel.Price
	hotelDto.IdMongo = hotel.IdMongo
	hotelDto.IdAmadeus = hotel.IdAmadeus
	return hotelDto, nil

}
