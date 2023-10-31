package services

import (
	bookingClient "uba/clients/booking"
	"uba/dto"
	"uba/model"

	//cache "uba/cache"
	//"time"
	e "uba/utils/errors"
)

type bookingService struct{}

type bookingServiceInterface interface {
	InsertBooking(bookingDto dto.BookingDto) (dto.BookingDto, e.ApiError)
}

var (
	BookingService bookingServiceInterface
)

func init() {
	BookingService = &bookingService{}
}

func (s *bookingService) InsertBooking(bookingDto dto.BookingDto) (dto.BookingDto, e.ApiError) {
	var booking model.Booking

	booking.Id = bookingDto.Id
	booking.StartDate = bookingDto.StartDate
	booking.EndDate = bookingDto.EndDate
	booking.IdUser = bookingDto.IdUser
	booking.IdHotel = bookingDto.IdHotel
	/*
		timeStart, _ := time.Parse("2023-05-20", bookingDto.StartDate)
		timeEnd, _ := time.Parse("2023-01-23", bookingDto.EndDate)

		var hotel model.Hotel
		var difference time.Duration
		var totalDays int32

		difference = timeEnd.Sub(timeStart)
		totalDays = int32(difference.Hours() / 24)
		booking.TotalPrice = float32(int32(hotel.Price) * totalDays)
	*/
	booking = bookingClient.InsertBooking(booking)

	var bookingResponseDto dto.BookingDto

	bookingResponseDto.Id = booking.Id
	bookingResponseDto.StartDate = booking.StartDate
	bookingResponseDto.EndDate = booking.EndDate
	bookingResponseDto.TotalPrice = booking.TotalPrice
	bookingResponseDto.IdUser = booking.IdUser
	bookingResponseDto.IdHotel = booking.IdHotel

	return bookingResponseDto, nil

}
