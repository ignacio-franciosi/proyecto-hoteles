package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"uba/clients/booking"
	"uba/clients/hotel"
	"uba/clients/user"
	"uba/dto"
	"uba/model"

	httpUtils "uba/utils/http"

	log "github.com/sirupsen/logrus"
)

type bookingService struct {
	HTTPClient httpUtils.HttpClientInterface
}

type bookingServiceInterface interface {
	InsertBooking(bookingDto dto.BookingDto) (dto.BookingDto, error)
	GetHotelInfo(idMongo string) (dto.HotelDto, error)
	GetAllHotelsByCity(city string) (dto.HotelsDto, error)
	CheckAvailability(idMongo string, startDate time.Time, endDate time.Time) bool
	CheckAllAvailability(city string, startDate string, endDate string) (dto.HotelsDto, error)
}

var (
	BookingService bookingServiceInterface
)

func init() {
	BookingService = &bookingService{
		HTTPClient: &httpUtils.HttpClient{},
	}
}

func (s *bookingService) InsertBooking(bookingDto dto.BookingDto) (dto.BookingDto, error) {

	userDto := user.GetUserById(bookingDto.IdUser)
	hotelDto, err := s.GetHotelInfo(bookingDto.IdMongo)

	if err != nil {
		return dto.BookingDto{}, errors.New("error retrieving hotel information")
	}

	if userDto.Id == 0 {
		return bookingDto, errors.New("user not found")
	}

	if hotelDto.IdMongo == "" {
		return bookingDto, errors.New("hotel not found")
	}

	// Parse dates with the correct format "DD-YY-MMMM"
	timeStart, err := time.Parse("02-01-2006", bookingDto.StartDate)
	if err != nil {
		return bookingDto, fmt.Errorf("invalid start date format: %v", err)
	}

	timeEnd, err := time.Parse("02-01-2006", bookingDto.EndDate)
	if err != nil {
		return bookingDto, fmt.Errorf("invalid end date format: %v", err)
	}

	if timeStart.After(timeEnd) {
		return bookingDto, errors.New("la reserva no puede terminar antes de empezar")
	}

	amadeusMap := hotel.GetHotelById(bookingDto.IdMongo)
	if amadeusMap.IdMongo == "" {
		return bookingDto, errors.New("no amadeus id set")
	}

	available, err := AmadeusService.CheckAvailabilityAmadeus(amadeusMap.IdAmadeus, timeStart, timeEnd)
	if err != nil {
		return bookingDto, err
	}

	if !available {
		return bookingDto, errors.New("no hay habitaciones (amadeus)")
	}

	if s.CheckAvailability(bookingDto.IdMongo, timeStart, timeEnd) {

		var booking2 model.Booking
		var diferencia time.Duration
		var cantDias int32

		booking2.StartDate = bookingDto.StartDate
		booking2.EndDate = bookingDto.EndDate
		booking2.Id = bookingDto.Id
		booking2.IdUser = bookingDto.IdUser
		booking2.IdMongo = bookingDto.IdMongo
		//------------------------------------------------------
		diferencia = timeEnd.Sub(timeStart)
		cantDias = int32(diferencia.Hours() / 24)
		booking2.TotalPrice = float64(int32(hotelDto.Price) * cantDias)

		fmt.Println("El valor de diferencia es:", diferencia)
		fmt.Println("El valor de cant dias es:", cantDias)
		fmt.Println("El valor de precio total es:", booking2.TotalPrice)

		booking2 = booking.InsertBooking(booking2)

		bookingDto.Id = booking2.Id
		bookingDto.TotalPrice = booking2.TotalPrice

		return bookingDto, nil
	}

	return bookingDto, errors.New("no hay habitaciones disponibles")
}

func (s *bookingService) GetHotelInfo(idMongo string) (dto.HotelDto, error) {
	resp, err := s.HTTPClient.Get("http://localhost:8080/hotel/" + idMongo)

	if err != nil {
		log.Error("Error in HTTP request ", err)
		return dto.HotelDto{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Error("Error reading response ", err)
		return dto.HotelDto{}, err
	}

	var hotelDto dto.HotelDto

	err = json.Unmarshal(body, &hotelDto)

	if err != nil {
		log.Error("Error parsing JSON ", err)
		return dto.HotelDto{}, err
	}

	return hotelDto, nil
}

func (s *bookingService) GetAllHotelsByCity(city string) (dto.HotelsDto, error) {

	var hotels model.Hotels = booking.GetAllHotelsByCity(city)
	var hotelsDto dto.HotelsDto

	for _, hotel := range hotels {
		var hotelDto dto.HotelDto

		hotelDto.Id = hotel.Id
		hotelDto.IdMongo = hotel.IdMongo
		hotelDto.IdAmadeus = hotel.IdAmadeus
		hotelDto.City = hotel.City
		hotelDto.Price = hotel.Price
		hotelDto.Rooms = hotel.Rooms

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}

func (s *bookingService) CheckAvailability(idMongo string, startDate time.Time, endDate time.Time) bool {
	fmt.Println("Entro al Check Availabilty")
	hotel, _ := s.GetHotelInfo(idMongo)
	fmt.Println("GetHotelInfo:", hotel)
	bookings := booking.GetBookingsByHotel(idMongo)
	fmt.Println("Las bookings: ", bookings)

	roomsAvailable := hotel.Rooms
	fmt.Println("hab disponibles segun la funcion:", roomsAvailable)

	for _, booking := range bookings {
		bookingStart, err := time.Parse("02-01-2006", booking.StartDate)
		if err != nil {
			// Log the error and continue to the next booking
			log.Errorf("Error parsing booking start date: %v", err)
			continue
		}

		bookingEnd, err := time.Parse("02-01-2006", booking.EndDate)
		if err != nil {
			// Log the error and continue to the next booking
			log.Errorf("Error parsing booking end date: %v", err)
			continue
		}

		// Check for overlapping bookings
		if (bookingStart.Before(endDate) && bookingEnd.After(startDate)) ||
			bookingStart.Equal(startDate) || bookingEnd.Equal(endDate) {
			roomsAvailable--
			fmt.Println("Se resto la habitacion")
		}

		if roomsAvailable == 0 {
			return false
		}
	}

	return true
}

func (s *bookingService) CheckAllAvailability(city string, startDate string, endDate string) (dto.HotelsDto, error) {
	//espera a que se completen las go routines
	var wg sync.WaitGroup
	var hotelsAvailable dto.HotelsDto

	bookingStart, _ := time.Parse("02-01-2006", startDate)
	bookingEnd, _ := time.Parse("02-01-2006", endDate)

	if bookingStart.After(bookingEnd) {
		return hotelsAvailable, errors.New("error, la reserva no puede terminar antes de comenzar")
	}

	//intenta obtener el resultado del cache. Parsea el resultado JSON y lo devuelve.

	fmt.Println("llego 1")

	hotels, err2 := s.GetAllHotelsByCity(city)
	if err2 != nil {
		// Manejar el error, por ejemplo, imprimirlo o devolverlo
		log.Println("Error al obtener los hoteles POR CIUDAD:", err2)
		return nil, err2
	}

	//canal para pasar los resultados de las goroutines que verifican la disponibilidad de los hoteles
	resultsCh := make(chan dto.HotelDto)

	for _, hotel := range hotels {
		//Esto incrementa el contador del WaitGroup en 1. Indica que hay una nueva goroutine que va a comenzar a ejecutar.
		wg.Add(1)
		go func(hotel dto.HotelDto) {
			//la goroutine decrementa el WaitGroup al finalizar con wg.Done()
			defer wg.Done()

			if s.CheckAvailability(hotel.IdMongo, bookingStart, bookingEnd) {
				var hotelDto dto.HotelDto
				hotelDto.Id = hotel.Id
				hotelDto.IdMongo = hotel.IdMongo
				hotelDto.IdAmadeus = hotel.IdAmadeus
				hotelDto.Rooms = hotel.Rooms
				hotelDto.Price = hotel.Price
				hotelDto.City = hotel.City

				resultsCh <- hotelDto
			}
		}(hotel)
	}
	//espera a que todas las goroutines de disponibilidad terminen (wg.Wait()) y luego cierra el canal resultsCh.
	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	for hotelDto := range resultsCh {
		hotelsAvailable = append(hotelsAvailable, hotelDto)
	}

	return hotelsAvailable, nil
}
