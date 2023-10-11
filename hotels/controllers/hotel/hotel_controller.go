package hotel

import (
	"hotels/config"
	"hotels/dto"
	service "hotels/services"
	client "hotels/services/repositories"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	// controlar
	hotelService = service.NewHotelServiceImpl(
		client.NewHotelInterface(config.MONGOHOST, config.MONGOPORT, config.MONGOCOLLECTION),
		client.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT),
	)
)

func GetHotelById(c *gin.Context) {
	var hotelDto dto.HotelDto
	id := c.Param("HotelId")
	hotelDto, err := hotelService.GetHotelById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, hotelDto)
}

func GetHotels(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	hotelsDto, err := hotelService.GetHotels()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, hotelsDto)

}

func InsertHotel(c *gin.Context) {
	var hotelDto dto.HotelDto
	err := c.BindJSON(&hotelDto)

	// Error Parsing json param
	if err != nil {

		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotelDto, er := hotelService.InsertHotel(hotelDto)

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}

func QueueHotels(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	err := c.BindJSON(&hotelsDto)

	// Error Parsing json param
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	er := hotelService.QueueHotels(hotelsDto)

	// Error Queueing
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelsDto)
}

func DeleteHotelById(c *gin.Context) {
	id := c.Param("HotelId")
	err := hotelService.DeleteHotelById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func UpdateHotelById(c *gin.Context) {

	var hotelDto dto.HotelDto
	hotelDto.HotelId = c.Param("HotelId")
	err := c.BindJSON(&hotelDto)

	// Error Parsing json param
	if err != nil {

		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotelDto, er := hotelService.UpdateHotelById(hotelDto)

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, hotelDto)
}
