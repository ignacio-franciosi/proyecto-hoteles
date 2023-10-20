package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	e "repo/utils/errors"
	"strings"
)

type bookingService struct{}

type bookingServiceInterface interface {
	InsertBooking(bookingDto dto.bookingDto) (dto.bookingDto, e.ApiError)
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

type Hotel struct {
	HotelId     string  `json:"hotel_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Amenities   string  `json:"amenities"`
	Stars       string  `json:"stars"`
	Rooms       int     `json:"rooms"`
	Price       float32 `json:"price"`
	City        string  `json:"city"`
	Photos      string  `json:"photos"`
}

func GetHotelsIds() ([]string, error) {
	// URL del microservicio que proporciona la función GetHoteles
	getHotelesURL := "http://direccion_del_microservicio/getHoteles"

	// Realizar la solicitud HTTP al microservicio
	resp, err := http.Get(getHotelesURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error al llamar al microservicio: %s", resp.Status)
	}

	// Leer la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Estructura para deserializar la respuesta
	var response []Hotel // Utilizando la estructura Hotel

	// Deserializar la respuesta JSON en la estructura
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// Extraer los HotelId
	var hotelIDs []string
	for _, hotel := range response {
		hotelIDs = append(hotelIDs, hotel.HotelId)
	}

	return hotelIDs, nil
}

func getHotelsAvailability() {
	// Define las variables de entrada
	hotelIds, err := GetHotelsIds()
	checkInDate := queryParams.Get("checkInDate")
	checkOutDate := "2024-01-24"

	// Define la URL de la llamada
	url := "https://test.api.amadeus.com/v3/shopping/hotel-offers?"

	// Agrega los parámetros a la URL
	url += "hotelIds=" + strings.Join(hotelIds, ",")
	url += "&checkInDate=" + checkInDate
	url += "&checkOutDate=" + checkOutDate

	// Define el token de autorización
	token, err := getBearer()

	// Crea una solicitud HTTP
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Agrega el encabezado de autorización
	req.Header.Set("Authorization", "Bearer "+token)

	// Realiza la solicitud
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Verifica el código de respuesta
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return
	}

	// Lee la respuesta
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Decodifica la respuesta
	offers := []Offer{}
	err = json.Unmarshal(bytes, &offers)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Filtra los hoteles con disponibilidad TRUE
	availableOffers := []Offer{}
	for _, offer := range offers {
		if offer.Availability == true {
			availableOffers = append(availableOffers, offer)
		}
	}

	// Imprime los hoteles con disponibilidad TRUE
	for _, offer := range availableOffers {
		fmt.Println(offer)
	}
}

func getBearer() (string, error) {
	// Datos de autenticación
	clientID := "8O9tG0NdPxxEA90mbxbAgRrJxGU002Yb"
	clientSecret := "qfurhUwj1OB0AjlE"

	// Crear la solicitud de token
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	// Realizar la solicitud HTTP para obtener el token de acceso
	resp, err := http.PostForm("https://test.api.amadeus.com/v3/security/oauth2/token", data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Verificar el código de estado de la respuesta
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error al obtener el token: código de estado %d", resp.StatusCode)
	}

	// Decodificar la respuesta JSON
	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return "", err
	}

	// Obtener el token de acceso
	accessToken, ok := responseBody["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("no se pudo obtener el token de acceso")
	}

	return accessToken, nil
}

/*
func GetHotelesDisponibles() []Hotel {
	// Obtiene la fecha desde, fecha hasta y ciudad
	hotels, err := GetHotelsFromHotels()
	fecha_desde := url.Query().Get("startDate")
	fecha_hasta := url.Query().Get("endDate")

	// Crea una estructura de datos para la solicitud a la API
	request := struct {
	  Hoteles []Hotel
	  FechaDesde string
	  FechaHasta string
	  Disponiblidad bool
	}{
	  Hoteles: hoteles,
	  FechaDesde: fecha_desde,
	  FechaHasta: fecha_hasta,
	  Disponiblidad: true,
	}

	// Realiza la llamada a la API
	response, err := http.Post("https://api.amadeus.com/v1/hotel/search", "application/json", json.Marshal(request))
	if err != nil {
	  return nil
	}

	// Decodifica la respuesta de la API
	var hoteles_disponibles []Hotel
	err = json.Unmarshal(response.Body, &hoteles_disponibles)
	if err != nil {
	  return nil
	}

	// Devuelve el slice de hoteles disponibles
	return hoteles_disponibles
  }


*/

/*
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
*/
