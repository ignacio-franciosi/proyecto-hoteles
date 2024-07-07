package services_test

import (
	"net/http"
	"testing"
	"time"
	"uba/dto"
	"uba/services"
	e "uba/utils/errors"

	"github.com/stretchr/testify/assert"
)

type TestBookings struct {
}

func (t *TestBookings) InsertBooking(bookingDto dto.BookingDto) (dto.BookingDto, error) {
	if bookingDto.Id == 0 {
		return dto.BookingDto{}, e.NewApiError("Error al insertar la reserva", "booking_insert_error", http.StatusInternalServerError, nil)
	}
	// Simulate successful insertion by setting a non-zero ID
	bookingDto.Id = 1
	return bookingDto, nil
}

func (t *TestBookings) GetAllHotelsByCity(city string) (dto.HotelsDto, error) {
	if city == "Miami" {
		return dto.HotelsDto{
			dto.HotelDto{
				Id:        1,
				IdMongo:   "1",
				IdAmadeus: "1",
				Rooms:     1,
				Price:     1,
				City:      "Miami",
			},
		}, nil
	}

	return dto.HotelsDto{}, e.NewNotFoundApiError("Hotel not found")
}

func (t *TestBookings) CheckAvailability(hotelId string, startDate time.Time, endDate time.Time) bool {
	// Aquí puedes agregar la lógica para determinar la disponibilidad
	return true
}

func (t *TestBookings) CheckAllAvailability(city string, startDate string, endDate string) (dto.HotelsDto, error) {
	// Aquí puedes agregar la lógica para verificar la disponibilidad de todos los hoteles en una ciudad
	return dto.HotelsDto{}, nil
}

func (t *TestBookings) GetHotelInfo(idMongo string) (dto.HotelDto, error) {
	if idMongo == "1" {
		return dto.HotelDto{
			Id:        1,
			IdMongo:   "1",
			IdAmadeus: "1",
			Rooms:     1,
			Price:     1,
			City:      "Miami",
		}, nil
	}
	return dto.HotelDto{}, e.NewNotFoundApiError("Hotel not found")
}

func TestInsertBooking(t *testing.T) {
	// Si cambio el valor de los id puedo ver los errores

	booking := dto.BookingDto{
		Id:         1,
		StartDate:  "2024-02-20",
		EndDate:    "2024-02-21",
		IdMongo:    "1",
		IdUser:     1,
		TotalPrice: 1.1,
	}

	services.BookingService = &TestBookings{}

	bookingCreated, err := services.BookingService.InsertBooking(booking)

	assert.Nil(t, err, "Error al insertar la reserva")
	assert.Equal(t, 1, bookingCreated.IdUser, "El ID de usuario no coincide")
	assert.Equal(t, "1", bookingCreated.IdMongo, "El ID de hotel no coincide")
	assert.Equal(t, booking.StartDate, bookingCreated.StartDate, "La fecha de inicio no coincide")
	assert.Equal(t, booking.EndDate, bookingCreated.EndDate, "La fecha de fin no coincide")
	assert.Equal(t, 1.1, bookingCreated.TotalPrice, "El precio total no coincide")

}

func TestGetHotelInfo(t *testing.T) {
	a := assert.New(t)

	services.BookingService = &TestBookings{}

	// Caso de prueba para un hotel existente
	idMongo := "1"
	hotel, err := services.BookingService.GetHotelInfo(idMongo)
	a.Nil(err)

	expectedHotel := dto.HotelDto{
		Id:        1,
		IdMongo:   "1",
		IdAmadeus: "1",
		Rooms:     1,
		Price:     1,
		City:      "Miami",
	}

	a.Equal(expectedHotel, hotel)

	// Caso de prueba para un hotel inexistente
	idMongo = "2"
	hotel, err = services.BookingService.GetHotelInfo(idMongo)
	a.NotNil(err)
	a.Contains(err.Error(), "Hotel not found")
	a.Equal(dto.HotelDto{}, hotel)
}

func TestGetAllHotelsByCity(t *testing.T) {
	// Caso de prueba para una reserva existente
	hotels, err := (&TestBookings{}).GetAllHotelsByCity("Miami")
	assert.Nil(t, err) // Asegúrate de que no haya errores

	// Asegúrate de que los valores de la reserva coincidan con lo esperado
	expectedHotel := dto.HotelDto{
		Id:        1,
		IdMongo:   "1",
		IdAmadeus: "1",
		Rooms:     1, // Ajusta según tus expectativas
		Price:     1,
		City:      "Miami",
	}

	assert.Equal(t, expectedHotel, hotels[0])

	// Caso de prueba para una reserva inexistente
	hotels, err = (&TestBookings{}).GetAllHotelsByCity("Madrid")
	assert.NotNil(t, err)                              // Asegúrate de que haya un error
	assert.Contains(t, err.Error(), "Hotel not found") // Asegúrate de que el mensaje de error contenga la cadena esperada
	assert.Empty(t, hotels)                            // Asegúrate de que no haya reservas en este caso
}

func TestCheckAvailability(t *testing.T) {
	a := assert.New(t)

	hotelId := "1"
	startDate := time.Date(2024, time.February, 20, 00, 00, 00, 00, time.UTC)
	endDate := time.Date(2024, time.February, 21, 00, 00, 00, 00, time.UTC)

	available := services.BookingService.CheckAvailability(hotelId, startDate, endDate)

	a.True(available)
}

func TestCheckAllAvailability(t *testing.T) {

	a := assert.New(t)

	city := "Miami"
	startDate1 := "20-02-2024"
	endDate1 := "21-02-2024"

	response1, err := services.BookingService.CheckAllAvailability(city, startDate1, endDate1)
	a.Nil(err)
	a.Equal(4, len(response1))

	startDate2 := "28-11-2023 15:00"
	endDate2 := "30-11-2023 11:00"

	response2, err := services.BookingService.CheckAllAvailability(city, startDate2, endDate2)
	a.Nil(err)
	a.Equal(5, len(response2))
}
