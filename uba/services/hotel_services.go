package services

import (
	"errors"
	hotelClient "uba/clients/hotel"
	"uba/dto"
	"uba/model"
	e "uba/utils/errors"
)

type hotelService struct{}

type hotelServiceInterface interface {
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
	GetHotelById(id string) (dto.HotelDto, e.ApiError)
	DeleteHotel(id string) error
	UpdateHotel(hotelDto dto.HotelDto) (dto.HotelDto, error)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {
	var hotel model.Hotel

	hotel.IdAmadeus = hotelDto.IdAmadeus
	hotel.IdMongo = hotelDto.IdMongo
	hotel.Rooms = hotelDto.Rooms
	hotel.Price = hotelDto.Price
	hotel.City = hotelDto.City

	hotel = hotelClient.InsertHotel(hotel)

	var response dto.HotelDto

	response.IdHotel = hotel.IdHotel
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

	hotelDto.IdHotel = hotel.IdHotel
	hotelDto.IdMongo = hotel.IdMongo
	hotelDto.IdAmadeus = hotel.IdAmadeus
	hotelDto.Rooms = hotel.Rooms
	hotelDto.Price = hotel.Price
	hotelDto.City = hotel.City
	return hotelDto, nil

}

func (s *hotelService) DeleteHotel(id string) error {

	hotel := hotelClient.GetHotelById(id)

	if hotel.IdMongo == "000000000000000000000000" {
		return errors.New("hotel not found")
	}

	err := hotelClient.DeleteHotel(hotel)

	return err
}

func (s *hotelService) UpdateHotel(hotelDto dto.HotelDto) (dto.HotelDto, error) {

	hotel := hotelClient.GetHotelById(hotelDto.IdMongo)

	if hotel.IdMongo == "000000000000000000000000" {
		return hotelDto, errors.New("hotel not found")
	}

	hotel.IdAmadeus = hotelDto.IdAmadeus
	hotel.IdMongo = hotelDto.IdMongo
	hotel.Rooms = hotelDto.Rooms
	hotel.Price = hotelDto.Price
	hotel.City = hotelDto.City

	hotel = hotelClient.UpdateHotel(hotel)

	if hotel.IdMongo == "000000000000000000000000" {
		return hotelDto, errors.New("error updating hotel")
	}

	return hotelDto, nil

}
