package hotel

import (
	"encoding/json"
	"fmt"
	"net/http"
	bookingController "uba/controllers/booking"
	"uba/dto"
	service "uba/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// esta funcion se llama cuando desde mongo se hace un post de hotel
func InsertHotel(c *gin.Context) {

	var insertHotelDto dto.InsertHotelDto
	err := c.BindJSON(&insertHotelDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	apiUrl := "https://test.api.amadeus.com/v1/reference-data/locations/hotels/by-city?cityCode=MIA&radius=5&radiusUnit=KM&hotelSource=ALL"
	// Crear una request HTTP GET
	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println("Error al crear la request:", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	token := bookingController.GetAmadeustoken()
	request.Header.Set("Authorization", "Bearer "+token)
	// Realiza la request HTTP
	cliente := &http.Client{}
	respuesta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("Error al realizar la request:", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	defer respuesta.Body.Close()

	// Leer la respuesta del amadeus
	var response struct {
		Data []struct {
			HotelID string `json:"hotelId"`
		} `json:"data"`
	}

	// Decodificar la respuesta JSON
	decoder := json.NewDecoder(respuesta.Body)
	if err := decoder.Decode(&response); err != nil {
		log.Error("Error al decodificar la respuesta JSON:", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	/*Este for recorre toda la estructura de Data hasta encontrar HotelId, cuando lo encuentra
	se lo pasa como parametro a la funcion*/
	for _, hotel := range response.Data {
		hotelDto, er := service.HotelService.InsertHotel(insertHotelDto, hotel.HotelID)
		// Error del Insert
		if er != nil {
			c.JSON(er.Status(), er)
			return
		}
		c.JSON(http.StatusCreated, hotelDto)
		log.Println("ID del hotel:", hotel.HotelID)
		break

	}

}
