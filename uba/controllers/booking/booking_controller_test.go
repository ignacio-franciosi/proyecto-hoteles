package booking_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	bookingController "uba/controllers/booking"
	"uba/dto"
	"uba/services"
	e "uba/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type TestBookings struct {
}

func (t *TestBookings) InsertBooking(bookingDto dto.BookingDto) (dto.BookingDto, error) {
	if bookingDto.IdUser == 0 {
		return dto.BookingDto{}, e.NewApiError("Error al insertar la reserva", "reserva_insert_error", http.StatusInternalServerError, nil)
	}

	return dto.BookingDto{}, nil
}

func (t *TestBookings) GetAllHotelsByCity(city string) (dto.HotelsDto, error) {
	if city == "Miami" {
		return dto.HotelsDto{
			dto.HotelDto{
				IdHotel:   1,
				IdMongo:   "1",
				IdAmadeus: "1",
				Rooms:     1,
				Price:     1,
				City:      "Miami",
			},
		}, nil
	}

	return dto.HotelsDto{}, e.NewNotFoundApiError("Booking not found")
}

func (t *TestBookings) CheckAvailability(_ string, _ time.Time, _ time.Time) bool {

	return true
}

func (t *TestBookings) CheckAllAvailability(city string, _ string, _ string) (dto.HotelsDto, error) {

	if city == "" {
		return dto.HotelsDto{}, errors.New("no city selected")
	}

	return dto.HotelsDto{}, nil
}

func (t *TestBookings) GetHotelInfo(hotelId string) (dto.HotelDto, error) {

	return dto.HotelDto{}, nil
}

func TestInsertBooking(t *testing.T) {
	services.BookingService = &TestBookings{}
	router := gin.Default()

	router.POST("/booking", bookingController.InsertBooking)

	// Solicitud HTTP POST - Si se cambia el User id a 0 se ve el error
	Json := `{
		"startDate":  "2024-02-20",
		"endDate":    "2024-02-21",
		"idMongo":    "1",
		"idUser":     1,
		"totalPrice": 1
		}`

	response := httptest.NewRecorder()

	bodyJson := strings.NewReader(Json)
	request, _ := http.NewRequest("POST", "/booking", bodyJson)

	router.ServeHTTP(response, request)

	fmt.Println(response.Body.String())

	// Verificar el código de estado de la respuesta
	assert.Equal(t, http.StatusCreated, response.Code, "El código de respuesta no es el esperado")

}

func TestGetAllHotelsByCity(t *testing.T) {
	services.BookingService = &TestBookings{}
	router := gin.Default()

	router.GET("/booking/:city", bookingController.GetAllHotelsByCity)

	// Crear una solicitud HTTP de tipo GET al endpoint /booking/{city}

	// Si se cambia el nombre de city por otra se ve el error
	request, _ := http.NewRequest("GET", "/booking/Miami", nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	fmt.Println(response.Body.String())

	// Verificar el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, response.Code, "ERROR DE CIUDAD")
}

func TestCheckAvailability(t *testing.T) {

	a := assert.New(t)

	r := gin.Default()
	r.GET("/available", bookingController.CheckAvailability)

	req, err := http.NewRequest(http.MethodGet, "/available?city=Miami", nil)
	if err != nil {
		t.Fatalf("New request failed: %v", err)
	}

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	r.ServeHTTP(w, req)

	a.Equal(http.StatusOK, w.Code)

	expectedResponse := dto.HotelsDto{}
	var response dto.HotelsDto

	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarhsal response: %v", err)
	}
	a.Equal(expectedResponse, response)
}
