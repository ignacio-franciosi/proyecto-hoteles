package hotel

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	bookingController "uba/controllers/booking"
	"uba/dto"
	service "uba/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CityCodeMapping struct {
	City string
	Code string
}

//Primero hago la funcion para obtener el city code que despues le tengo que pasar al url de amadeus

func GetCityCode() string {

	var insertHotelDto dto.InsertHotelDto
	cityName := insertHotelDto.City

	// slice de mapeo de City a código
	mapCityCode := []CityCodeMapping{
		{City: "Barcelona", Code: "BCN"},
		{City: "New York", Code: "NYC"},
		{City: "Paris", Code: "PAR"},
		{City: "Milan", Code: "MIL"},
		{City: "Toronto", Code: "YTO"},
		{City: "Sydney", Code: "SYD"},
		{City: "Seoul", Code: "SEL"},
		{City: "Moscow", Code: "MOW"},
		{City: "London", Code: "LON"},
	}

	for _, mapp := range mapCityCode {
		if mapp.City == cityName {
			return mapp.Code
		}
	}

	return ""
}

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

	apiUrl := "https://test.api.amadeus.com/v1/reference-data/locations/hotels/by-city"
	apiUrl += "?cityCode=" + GetCityCode()
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

func GetHotelById(c *gin.Context) {

	log.Debug("Hotel id: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var hotelDto dto.HotelDto

	hotelDto, err := service.HotelService.GetHotelById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}
