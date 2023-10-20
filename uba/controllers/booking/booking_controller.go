package booking

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"uba/dto"
	service "uba/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// funcion que genera el token de amadeus
func GetAmadeustoken() string {

	// credenciales
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", "8O9tG0NdPxxEA90mbxbAgRrJxGU002Yb")
	data.Set("client_secret", "qfurhUwj1OB0AjlE")

	// obtengo el token
	resp, err := http.Post("https://test.api.amadeus.com/v3/security/oauth2/token", "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer resp.Body.Close()
	// leo la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	// Decodifico el JSON
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return ""
	}
	// Obtengo el token
	token := response["access_token"].(string)

	fmt.Println("token:", token)

	return token

}

func GetAvailabilityByIdAndDate(c *gin.Context) {
	var bookingDto dto.BookingDto

	err := c.BindJSON(&bookingDto)
	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// ahora con los datos del booking dto rellenamos una nueva estructura para la requesta  amadeus
	// Serializa el objeto BookingDto a formato JSON
	id := bookingDto.HotelId
	fmt.Println("El id mysql del hotel es:", id)
	// necesito llmara a una funcion que me traiga el id amadeus del hotel con el id que ya tengo (tengo el id mysql)
	// GetHotelById(id int) (dto.HotelDto, e.ApiError)
	hotelDto, err := service.HotelService.GetHotelById(id)

	if err != nil {
		// c.JSON(err.Status(), err)
		fmt.Println("No se encontro un hotel con ese id")
		return
	}

	idAm := hotelDto.IdAmadeus

	startdatebooking := strconv.Itoa(bookingDto.StartDate)
	fechaConGuiones := startdatebooking
	startdateconguiones := fmt.Sprintf(
		"%s-%s-%s",
		fechaConGuiones[:4],
		fechaConGuiones[4:6],
		fechaConGuiones[6:8],
	)
	enddatebooking := strconv.Itoa(bookingDto.EndDate)
	fechaConGuiones2 := enddatebooking
	enddateconguiones := fmt.Sprintf(
		"%s-%s-%s",
		fechaConGuiones2[:4],
		fechaConGuiones2[4:6],
		fechaConGuiones2[6:8],
	)

	fmt.Println("fecha de ida", startdateconguiones)
	fmt.Println("fecha de vuelta", enddateconguiones)

	// Construye la URL manualmente
	apiUrl := "https://test.api.amadeus.com/v3/shopping/hotel-offers"
	apiUrl += "?hotelIds=" + idAm
	apiUrl += "&checkInDate=" + startdateconguiones
	apiUrl += "&checkOutDate=" + enddateconguiones

	fmt.Println(apiUrl)

	// Crear una solicitud HTTP GET
	solicitud, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println("Error al crear la solicitud:", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Agregar el encabezado de autorización Bearer con tu token
	token := GetAmadeustoken() // Reemplaza con tu token real

	solicitud.Header.Set("Authorization", "Bearer "+token)
	// solicitud.Header.Set("Content-Type", "application/json") // Especifica el tipo de contenido si es necesario

	fmt.Println(solicitud)
	// Realiza la solicitud HTTP
	cliente := &http.Client{}
	respuesta, err := cliente.Do(solicitud)
	if err != nil {
		fmt.Println("Error al realizar la solicitud:", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else if err == nil {
		// Verifica el código de estado de la respuesta
		if respuesta.StatusCode != http.StatusOK {
			fmt.Printf("La solicitud a la API de Amadeus no fue exitosa. Código de estado: %d\n", respuesta.StatusCode)
			c.JSON(http.StatusInternalServerError, "La solicitud a la API de Amadeus no fue exitosa.")
			return
		}
		// Lee el cuerpo de la respuesta
		responseBody, err := ioutil.ReadAll(respuesta.Body)
		if err != nil {
			fmt.Println("Error al leer la respuesta:", err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		// Crear una estructura para deserializar el JSON de la respuesta
		var responseStruct struct {
			Data []struct {
				Type                   string `json:"type"`
				ID                     string `json:"id"`
				ProviderConfirmationID string `json:"providerConfirmationId"`
			} `json:"data"`
		}

		// Decodificar el JSON y extraer el campo "id"
		if err := json.Unmarshal(responseBody, &responseStruct); err != nil {
			fmt.Println("Error al decodificar el JSON de la respuesta:", err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		// Obtén el ID del hotel del primer elemento en "data"
		if len(responseStruct.Data) > 0 {
			// si el largo de la respuesta es mayor q cero es pq hay disponibilidad --> llamo al service

			bookingDto, err := service.BookingService.InsertBooking(bookingDto)
			// Error del Insert
			if err != nil {
				// c.JSON(err.Status(), err)
				return
			}
			c.JSON(http.StatusCreated, bookingDto)
		} else {
			fmt.Println("No hay disponibilidad en esas fechas")
		}

		defer respuesta.Body.Close()

	}
}

func InsertBooking(c *gin.Context) {
	var bookingDto dto.BookingDto
	err := c.BindJSON(&bookingDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	bookingDto, er := service.BookingService.InsertBooking(bookingDto)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, bookingDto)

}
