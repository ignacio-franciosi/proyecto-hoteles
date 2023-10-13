package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	bookingClient "repo/clients/booking"
	"repo/dto"
	model "repo/model"
	e "repo/utils/errors"
	"time"
	"uba/dto"
)

type bookingService struct{}

type bookingServiceInterface interface {
	InsertBooking(bookingDto dto.bookingDto) (dto.bookingDto, e.ApiError)
	GetbookingsByIdUser(token string) (dto.bookingsDto, e.ApiError) // ----------------
}

var (
	BookingService bookingServiceInterface
)

func init() {
	BookingService = &bookingService{}
}

type AmadeusCredentials struct {
	ClientID     string
	ClientSecret string
}

func getAccessToken() (string, error) {

	credentials := AmadeusCredentials{
		ClientID:     "8O9tG0NdPxxEA90mbxbAgRrJxGU002Yb",
		ClientSecret: "qfurhUwj1OB0AjlE",
	}

	u, err := url.Parse("https://test.api.amadeus.com/v3/security/oauth2/token")
	if err != nil {
		return "", err
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", credentials.ClientID)
	data.Set("client_secret", credentials.ClientSecret)

	resp, err := http.PostForm(u.String(), data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get access token, status code: %d", resp.StatusCode)
	}

	var tokenResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		return "", err
	}

	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("failed to parse access token from response")
	}

	return accessToken, nil
}

func getHotelAvailability(dto.HotelsDto) (dto.HotelsDto, error) {
	// Endpoint para obtener disponibilidad de hoteles
	baseURL := "https://test.api.amadeus.com/v3/shopping/hotel-offers"
	var booking model.Booking
	var hotelesDto dto.HotelesDto

	// Obtén el token de acceso
	accessToken, err := getAccessToken()
	if err != nil {
		return hotelesDto, err
	}

	for _, hotel := range hotels {
		queryParams := url.Values{}
		queryParams.Set("city", hotel.City)
		queryParams.Set("startDate", booking.StartDate) // Ejemplo de fecha de inicio
		queryParams.Set("endDate", booking.EndDate)     // Ejemplo de fecha de fin

		fullURL := baseURL + "?" + queryParams.Encode()

		// Crea una solicitud HTTP con el método GET
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return hotelesDto, err
		}

		// Agrega el token de acceso en el encabezado de autorización
		req.Header.Set("Authorization", "Bearer "+accessToken)

		// Realiza la solicitud
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return hotelesDto, err
		}
		defer resp.Body.Close()

		// Leer y procesar la respuesta aquí según tus necesidades
		// Puedes leer el cuerpo de la respuesta usando resp.Body

		// Agregar la lógica para procesar la respuesta y agregar a hotelesDto
		// hotelesDto.AgregarHotel(...) - Agrega un hotel a la estructura de respuesta
	}

	return hotelesDto, nil
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

	url := fmt.Sprintf("http://microservicio-hoteles.com/getHotelById?id=%s", idHotel)

	hotel = hotelCliente.GetHotelById(booking.IdHotel) //---------

	difference = booking.EndDate.Sub(booking.StartDate)
	totalDays = int32(difference.Hours() / 24)
	booking.TotalPrice = float32(int32(hotel.Precio) * totalDays) // -------
	booking = bookingClient.InsertBooking(booking)

	var bookingResponseDto dto.bookingDto

	bookingResponseDto.Id = booking.Id
	bookingResponseDto.StartDate = booking.StartDate
	bookingResponseDto.EndDate = booking.EndDate
	bookingResponseDto.TotalPrice = booking.TotalPrice
	bookingResponseDto.IdUser = booking.IdUser
	bookingResponseDto.IdHotel = booking.IdHotel

	return bookingResponseDto, nil

}
