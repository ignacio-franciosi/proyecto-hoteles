package booking

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	data.Set("client_id", "rZ5NyDeG74mzhfaFqk2V7VaM2yUQtZux")
	data.Set("client_secret", "IoO2Wocmr1Aqrw2X")

	// obtengo el token
	resp, err := http.Post("https://test.api.amadeus.com/v3/security/oauth2/token", "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer resp.Body.Close()
	// leo la response
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

func InsertBooking(c *gin.Context) {

	var bookingDto dto.BookingDto

	err := c.BindJSON(&bookingDto)
	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// Serializo el BookingDto a formato JSON

	id := bookingDto.IdHotel
	fmt.Println("El id mysql del hotel es:", id)
	// necesito llmara a una funcion que me traiga el id amadeus del hotel con el id que ya tengo (tengo el id mysql)

	hotelDto, err := service.HotelService.GetHotelById(id)

	if err != nil {
		// c.JSON(err.Status(), err)
		fmt.Println("No se encontro un hotel con ese id")
		return
	}

	idAmad := hotelDto.IdAmadeus
	startDateBooking := bookingDto.StartDate
	endDateBooking := bookingDto.EndDate

	fmt.Println("fecha de ida", startDateBooking)
	fmt.Println("fecha de vuelta")

	// Construyo la URL concatenando los parametros
	apiUrl := "https://test.api.amadeus.com/v3/shopping/hotel-offers"
	apiUrl += "?hotelIds=" + idAmad
	apiUrl += "&checkInDate=" + startDateBooking
	apiUrl += "&checkOutDate=" + endDateBooking

	fmt.Println(apiUrl)

	// Hago la request HTTP GET
	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println("Error al crear la request:", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Llamo a la funcion que obtiene el token
	token := GetAmadeustoken()

	request.Header.Set("Authorization", "Bearer "+token)
	// request.Header.Set("Content-Type", "application/json")

	fmt.Println(request)
	// Realiza la request HTTP
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error al realizar la request:", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	} else if err == nil {
		// Verifica el código de estado de la response
		if response.StatusCode != http.StatusOK {
			fmt.Printf("La request a la API de Amadeus no fue exitosa. Código: %d\n", response.StatusCode)
			c.JSON(http.StatusInternalServerError, "La request a la API de Amadeus no fue exitosa.")
			return
		}
		// Lee el cuerpo de la response
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error al leer la response:", err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		// Crear una estructura para deserializar el JSON de la response
		var responseStruct struct {
			Data []struct {
				HotelId string `json:"hotelId"`
			} `json:"data"`
		}

		// Decodificar el JSON y extraer el campo "id"
		if err := json.Unmarshal(responseBody, &responseStruct); err != nil {
			fmt.Println("Error al decodificar el JSON de la response:", err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		// Obtén el ID del hotel del primer elemento en "data"
		if len(responseStruct.Data) > 0 {
			// si la response no esta vacia, es porque algo me devolvio o sea si esta disponible

			bookingDto, err := service.BookingService.InsertBooking(bookingDto)
			// Error del Insert
			if err != nil {
				// c.JSON(err.Status(), err)
				return
			}
			c.JSON(http.StatusCreated, bookingDto)
		} else {
			fmt.Println("No hay disponibilidad en ese periodo")
		}

		defer response.Body.Close()

	}
}
