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
