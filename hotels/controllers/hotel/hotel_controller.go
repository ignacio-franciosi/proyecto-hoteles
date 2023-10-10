package hotel

import (
	"hotels/config"
	"hotels/dto"
	service "hotels/services"
	client "hotels/services/repositories"
	
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	// controlar
	hotelService = service.NewHotelServiceImpl(
		client.NewHotelInterface(config.MONGOHOST, config.MONGOPORT, config.MONGOCOLLECTION),
		client.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT),
	)
)

func GetHotelById(c *gin.Context) {
	var hotelDto dto.HotelResponseDto
	id := c.Param("HotelId")
	hotelDto, err := hotelService.GetHotelById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, hotelDto)
}

func GetItemsByUserId(c *gin.Context) {
	var itemsDto dto.ItemsResponseDto
	id, _ := strconv.Atoi(c.Param("id"))
	itemsDto, err := itemService.GetItemsByUserId(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, itemsDto)

}

func InsertItem(c *gin.Context) {
	var itemDto dto.ItemDto
	err := c.BindJSON(&itemDto)

	// Error Parsing json param
	if err != nil {

		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	itemDto, er := itemService.InsertItem(itemDto)

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, itemDto)
}

func QueueItems(c *gin.Context) {
	var itemsDto dto.ItemsDto
	err := c.BindJSON(&itemsDto)

	// Error Parsing json param
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	er := itemService.QueueItems(itemsDto)

	// Error Queueing
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, itemsDto)
}

func DeleteUserItems(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := itemService.DeleteUserItems(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func DeleteItemById(c *gin.Context) {
	id := c.Param("item_id")
	err := itemService.DeleteItemById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}