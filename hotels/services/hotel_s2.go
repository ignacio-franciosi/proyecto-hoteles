package services

import (
	"encoding/json"
	"errors"
	"hotels/client"
	"hotels/dto"
	"hotels/model"
	"hotels/queue"
	"hotels/utils"
	"net/http"
)

type hotelService struct {
	HTTPClient utils.HttpClientInterface
}

type hotelServiceInterface interface {
	GetHotelById(id string) (dto.HotelDto, error)
	GetAllHotels() (dto.HotelsDto, error)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error)
	DeleteHotel(id string) (dto.HotelDto, error)
	UpdateHotel(hotelDto dto.HotelDto) (dto.HotelDto, error)
}

var HotelService hotelServiceInterface

func init() {
	HotelService = &hotelService{
		HTTPClient: &utils.HttpClient{},
	}
}

func (s *hotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, error) {

	var hotel model.Hotel

	hotel.Name = hotelDto.Name
	hotel.Description = hotelDto.Description
	hotel.Amenities = hotelDto.Amenities
	hotel.Stars = hotelDto.Stars
	hotel.Rooms = hotelDto.Rooms
	hotel.Price = hotelDto.Price
	hotel.City = hotelDto.City
	hotel.Photos = hotelDto.Photos

	hotel = client.HotelClient.InsertHotel(hotel)

	hotelDto.HotelId = hotel.HotelId.Hex()

	if hotel.HotelId.Hex() == "000000000000000000000000" {
		return hotelDto, errors.New("error creating hotel")
	}

	body := map[string]interface{}{
		"Id":      hotel.HotelId.Hex(),
		"Message": "create",
	}

	jsonBody, _ := json.Marshal(body)

	err := queue.QueueProducer.Publish(jsonBody)

	if err != nil {
		return hotelDto, err
	}

	return hotelDto, nil
}

func (s *hotelService) GetAllHotels() (dto.HotelsDto, error) {

	var hotels model.Hotels = client.HotelClient.GetAllHotels()
	var hotelsDto dto.HotelsDto

	for _, hotel := range hotels {
		var hotelDto dto.HotelDto
		hotelDto.HotelId = hotel.HotelId.Hex()
		hotelDto.Name = hotel.Name
		hotelDto.Description = hotel.Description
		hotelDto.Stars = hotel.Stars
		hotelDto.Rooms = hotel.Rooms
		hotelDto.Price = hotel.Price
		hotelDto.City = hotel.City
		hotelDto.Photos = hotel.Photos

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}

func (s *hotelService) GetHotelById(id string) (dto.HotelDto, error) {

	var hotelDto dto.HotelDto

	hotel, err := client.HotelClient.GetHotelById(id)
	/*
		if hotel.HotelId.Hex() == "000000000000000000000000" {
			return hotelDto, errors.New("hotel not found")
		}
	*/
	if err != nil {
		return hotelDto, errors.New("hotel not found")
	}

	hotelDto.HotelId = hotel.HotelId.Hex()
	hotelDto.Name = hotel.Name
	hotelDto.Description = hotel.Description
	hotelDto.Stars = hotel.Stars
	hotelDto.Rooms = hotel.Rooms
	hotelDto.Price = hotel.Price
	hotelDto.City = hotel.City
	hotelDto.Photos = hotel.Photos

	return hotelDto, nil
}

func (s *hotelService) DeleteHotel(id string) (dto.HotelDto, error) {

	var hotelDto dto.HotelDto

	hotel, err := client.HotelClient.GetHotelById(id)
	if err != nil {
		return hotelDto, errors.New("hotel not found")
	}

	if hotel.HotelId.Hex() == "000000000000000000000000" {
		return dto.HotelDto{}, errors.New("hotel not found")
	}

	url := "http://user-reservationnginx:8090/amadeus/" + id
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return dto.HotelDto{}, err
	}

	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return dto.HotelDto{}, err
	}

	if res.StatusCode != http.StatusOK {
		return dto.HotelDto{}, errors.New("status not OK")
	}

	err = client.HotelClient.DeleteHotelById(id)

	if err != nil {
		return dto.HotelDto{}, err
	}

	body := map[string]interface{}{
		"HoteId":  hotel.HotelId.Hex(),
		"message": "delete",
	}

	jsonBody, _ := json.Marshal(body)

	err = queue.QueueProducer.Publish(jsonBody)

	if err != nil {
		return dto.HotelDto{}, err
	}

	hotelDto.HotelId = hotel.HotelId.Hex()

	return hotelDto, err
}

func (s *hotelService) UpdateHotel(hotelDto dto.HotelDto) (dto.HotelDto, error) {

	hotel, err := client.HotelClient.GetHotelById(hotelDto.HotelId)
	if err != nil {
		return hotelDto, errors.New("hotel not found")
	}

	if hotel.HotelId.Hex() == "000000000000000000000000" {
		return hotelDto, errors.New("hotel not found")
	}

	hotel.Name = hotelDto.Name
	hotel.Description = hotelDto.Description
	hotel.Amenities = hotelDto.Amenities
	hotel.Stars = hotelDto.Stars
	hotel.Rooms = hotelDto.Rooms
	hotel.Price = hotelDto.Price
	hotel.City = hotelDto.City
	hotel.Photos = hotelDto.Photos

	updateHotel := client.HotelClient.UpdateHotelById(hotelDto.HotelId, hotel)

	if updateHotel.HotelId.Hex() == "000000000000000000000000" {
		return hotelDto, errors.New("error updating hotel")
	}

	body := map[string]interface{}{
		"HotelId": hotel.HotelId.Hex(),
		"message": "update",
	}

	jsonBody, _ := json.Marshal(body)

	err = queue.QueueProducer.Publish(jsonBody)

	if err != nil {
		return hotelDto, err
	}

	return hotelDto, nil

}
