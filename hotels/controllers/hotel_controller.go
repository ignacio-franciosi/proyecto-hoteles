package controllers

import (
	"hotels/dto"
	service "hotels/services"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InsertHotel(c *gin.Context) {
	var hotelDto dto.HotelDto
	err := c.BindJSON(&hotelDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotelDto, er := service.HotelService.InsertHotel(hotelDto)

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}

func GetHotelById(c *gin.Context) {

	id := c.Param("HotelId")
	var hotelDto dto.HotelDto

	hotelDto, err := service.HotelService.GetHotelById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}

func GetAllHotels(c *gin.Context) {

	var hotelsDto dto.HotelsDto

	hotelsDto, err := service.HotelService.GetAllHotels()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hotelsDto)
}

func DeleteHotel(c *gin.Context) {
	id := c.Param("HotelId")

	hotel, err := service.HotelService.DeleteHotel(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hotel " + hotel.HotelId + " deleted"})
}

func UpdateHotel(c *gin.Context) {
	id := c.Param("HotelId")
	var hotelDto dto.HotelDto
	err := c.BindJSON(&hotelDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hotelDto.HotelId = id

	hotelDto, err = service.HotelService.UpdateHotel(hotelDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hotelDto)
}
