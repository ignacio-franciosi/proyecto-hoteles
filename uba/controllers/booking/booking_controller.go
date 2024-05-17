package booking

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
	"uba/dto"
	service "uba/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

var (
	amadeusToken    string
	mutex           sync.Mutex
	lastTokenUpdate time.Time
)

func GetAmadeusToken() (string, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Verificar si el token necesita actualización (más de 30 minutos desde la última actualización)
	if time.Since(lastTokenUpdate) > 30*time.Minute {
		err := updateAmadeusToken()
		if err != nil {
			return "", err
		}
	}

	// Devolver el token actual
	return amadeusToken, nil
}

func updateAmadeusToken() error {
	url := "https://test.api.amadeus.com/v1/security/oauth2/token"
	method := "POST"

	payload := strings.NewReader("grant_type=client_credentials&client_id=rZ5NyDeG74mzhfaFqk2V7VaM2yUQtZux&client_secret=IoO2Wocmr1Aqrw2X")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var response TokenResponse

	// Decodificar la respuesta JSON en la estructura TokenResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	// Actualizar el token y la hora de la última actualización
	amadeusToken = response.AccessToken
	lastTokenUpdate = time.Now()

	fmt.Println("Token actualizado:", amadeusToken)
	return nil
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

	id := bookingDto.IdMongo
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
	fmt.Println("fecha de vuelta", endDateBooking)

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
	token, er := GetAmadeusToken()
	if er != nil {
		fmt.Println("Error:", er)
		return
	}
	fmt.Println(token)

	request.Header.Set("Authorization", "Bearer "+token)

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
			fmt.Printf("No hay disponibilidad. Código: %d\n", response.StatusCode)
			c.JSON(http.StatusInternalServerError, "No hay disponibilidad.")
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
			Data []map[string]interface{} `json:"data"`
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
			errorMessage := "No hay disponibilidad en ese período"
			c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
			return
		}

		defer response.Body.Close()

	}
}
