package hotel

import (
	"net/http"
	"uba/dto"
	service "uba/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//Primero hago la funcion para obtener el city code que despues le tengo que pasar al url de amadeus

// esta funcion se llama cuando desde mongo se hace un post de hotel
func InsertHotel(c *gin.Context) {

	var hotelDto dto.HotelDto
	err := c.BindJSON(&hotelDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotelDto, er := service.HotelService.InsertHotel(hotelDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	c.JSON(http.StatusCreated, hotelDto)

}

func GetHotelById(c *gin.Context) {

	log.Debug("Hotel id: " + c.Param("id"))

	id := (c.Param("id"))
	var hotelDto dto.HotelDto

	hotelDto, err := service.HotelService.GetHotelById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}

func DeleteHotel(c *gin.Context) {
	id := c.Param("id")

	err := service.HotelService.DeleteHotel(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hotel deleted"})
}

func UpdateHotel(c *gin.Context) {
	id := c.Param("id")
	var hotelDto dto.HotelDto
	err := c.BindJSON(&hotelDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotelDto.IdMongo = id

	hotelDto, err = service.HotelService.UpdateHotel(hotelDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hotelDto)
}
