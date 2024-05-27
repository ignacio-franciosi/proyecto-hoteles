package booking

import (
	"fmt"
	"net/http"
	"strings"
	"uba/dto"
	service "uba/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InsertBooking(c *gin.Context) {
	var bookingDto dto.BookingDto
	err := c.BindJSON(&bookingDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookingDto, er := service.BookingService.InsertBooking(bookingDto)

	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": er.Error()})
		return
	}

	c.JSON(http.StatusCreated, bookingDto)
}

func CheckAvailability(c *gin.Context) {

	city := c.Query("city")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	fmt.Println("Los valores que entran al checkALL:", city, startDate, endDate)
	hotelsAvailable, err := service.BookingService.CheckAllAvailability(city, startDate, endDate)
	fmt.Println("hotelsAv", hotelsAvailable)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, hotelsAvailable)
}

func GetAllHotelsByCity(c *gin.Context) {

	var hotelsDto dto.HotelsDto
	city := c.Param("city")
	cityFormatted := strings.ReplaceAll(city, " ", "+")
	hotelsDto, err := service.BookingService.GetAllHotelsByCity(cityFormatted)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, hotelsDto)

}
