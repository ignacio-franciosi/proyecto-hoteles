package booking_test

import (
	"testing"
	"uba/model"

	"github.com/stretchr/testify/assert"
)

type MockBookingClient struct{}

func (m *MockBookingClient) InsertBooking(booking model.Booking) model.Booking {
	// Simular la lógica de inserción en la base de datos
	// Se establece un ID para la booking
	booking.Id = 1 // Si modifico a cero, genera la alerta
	return booking
}

func (m *MockBookingClient) GetBookingsByHotel(idMongo string) model.Booking {
	// Simular la búsqueda en la base de datos
	booking := model.Booking{
		Id:         1,
		IdUser:     1,
		IdMongo:    "1",
		StartDate:  "20-02-2024",
		EndDate:    "21-02-2024",
		TotalPrice: 1111,
	}

	return booking
}

func (m *MockBookingClient) GetAllHotelsByCity(city string) model.Hotel {

	hotel := model.Hotel{
		Id:        1,
		IdMongo:   "1",
		IdAmadeus: "1",
		Rooms:     1,
		Price:     1,
		City:      "Miami",
	}

	return hotel

}

func TestInsertBooking(t *testing.T) {
	// Crear una instancia del mock del DAO de booking
	mockClient := &MockBookingClient{}

	// Crear una nueva booking
	newBooking := model.Booking{
		IdUser:     1,
		IdMongo:    "1",
		StartDate:  "20-02-2024",
		EndDate:    "21-02-2024",
		TotalPrice: 123,
	}

	// Insertar la booking utilizando el mock del DAO
	insert := mockClient.InsertBooking(newBooking)

	// Verificar que la booking tenga un ID asignado
	assert.NotZero(t, insert.Id, "La booking no se pudo realizar")

	// Verificar otros atributos de la booking
	assert.Equal(t, newBooking.IdUser, insert.IdUser)
	assert.Equal(t, newBooking.IdMongo, insert.IdMongo)
	assert.Equal(t, newBooking.StartDate, insert.StartDate)
	assert.Equal(t, newBooking.EndDate, insert.EndDate)
	assert.Equal(t, newBooking.TotalPrice, insert.TotalPrice)
}

func TestGetBookingsByHotel(t *testing.T) {
	// Crear una instancia del mock del DAO de booking
	mockClient := &MockBookingClient{}

	// ID de booking a buscar - Si la cambio deja de funcionar
	hotelId := "1"

	// Obtener la booking utilizando el mock del DAO
	booking := mockClient.GetBookingsByHotel(hotelId)

	// Verificar que la booking obtenida tenga el ID correcto
	assert.Equal(t, hotelId, booking.IdMongo, "El ID del hotel no existe")

	// Verificar otros atributos de la booking
	assert.Equal(t, 1, booking.IdUser)
	expectedDateFrom := "20-02-2024"
	assert.Equal(t, expectedDateFrom, booking.StartDate)
	expectedDateTo := "21-02-2024"
	assert.Equal(t, expectedDateTo, booking.EndDate)
}

func TestGetAllHotelsByCity(t *testing.T) {
	// Crear una instancia del mock del DAO de booking
	mockClient := &MockBookingClient{}

	// ID de booking a buscar - Si la cambio deja de funcionar
	city := "Miami"

	// Obtener la booking utilizando el mock del DAO
	hotel := mockClient.GetAllHotelsByCity(city)

	// Verificar que la booking obtenida tenga el ID correcto
	assert.Equal(t, city, hotel.City, "No coincide")

	// Verificar otros atributos de la booking
	assert.Equal(t, "1", hotel.IdMongo)
	assert.Equal(t, "1", hotel.IdAmadeus)
	assert.Equal(t, 1, hotel.Rooms)

}
