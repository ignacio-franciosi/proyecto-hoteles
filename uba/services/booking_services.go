package service

import (
	bookingClient "uba/clients/booking"
	"uba/dto"
	"uba/model"

	//cache "uba/cache"
	"time"
	e "uba/utils/errors"
)

type bookingService struct{}

type bookingServiceInterface interface {
	InsertBooking(bookingDto dto.BookingDto) (dto.BookingDto, e.ApiError)
}

func (s *bookingService) InsertBooking(bookingDto dto.BookingDto) (dto.BookingDto, e.ApiError) {
	var booking model.Booking

	booking.Id = bookingDto.Id
	booking.StartDate = bookingDto.StartDate
	booking.EndDate = bookingDto.EndDate
	booking.IdUser = bookingDto.IdUser
	booking.IdHotel = bookingDto.IdHotel

	var hotel model.Hotel
	var difference time.Duration
	var totalDays int32

	difference = booking.EndDate.Sub(booking.StartDate)
	totalDays = int32(difference.Hours() / 24)
	booking.TotalPrice = float32(int32(hotel.Price) * totalDays)
	booking = bookingClient.InsertBooking(booking)

	var bookingResponseDto dto.BookingDtoDto

	bookingResponseDto.Id = booking.Id
	bookingResponseDto.StartDate = booking.StartDate
	bookingResponseDto.EndDate = booking.EndDate
	bookingResponseDto.TotalPrice = booking.TotalPrice
	bookingResponseDto.IdUser = booking.IdUser
	bookingResponseDto.IdHotel = booking.IdHotel

	return bookingResponseDto, nil

}
