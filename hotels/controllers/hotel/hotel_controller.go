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
	hotelService = service.NewHotelServiceImpl(
		client.NewHotelInterface(config.MONGOHOST, config.MONGOPORT, config.MONGOCOLLECTION),
		client.NewQueueClient(config.RABBITUSER, config.RABBITPASSWORD, config.RABBITHOST, config.RABBITPORT),
	)
)

// Maneja las solicitudes para obtener un hotel por su ID. Llama al service para
// recuperar la información del hotel y envía una respuesta JSON al client.
func GetHotelById(c *gin.Context) {
	//var hotelDto dto.HotelDto
	id := c.Param("HotelId")
	hotelResponseDto, err := hotelService.GetHotelById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, hotelResponseDto)
}

// Maneja las solicitudes para obtener la lista de hoteles. Llama al service
// para obtener los datos y envía una respuesta al client con la lista de hoteles
func GetHotels(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	hotelsDto, err := hotelService.GetHotels()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, hotelsDto)

}

// Maneja las solicitudes para Insertar un nuevo hotel utilizando el service
// y luego envía una respuesta JSON al client
func InsertHotel(c *gin.Context) {
	var hotelDto dto.HotelDto
	err := c.BindJSON(&hotelDto)

	if err != nil {

		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotelDto, er := hotelService.InsertHotel(hotelDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}

// Maneja las solicitudes para encolar hoteles, los encola utilizando
// el service y luego envía una respuesta JSON al cliente,
func QueueHotels(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	err := c.BindJSON(&hotelsDto)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	er := hotelService.QueueHotels(hotelsDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelsDto)
}

// Maneja las solicitudes para eliminar un hotel por su ID obtenido por la URL
// llama al service para realizar la eliminación y envía una respuesta al client
func DeleteHotelById(c *gin.Context) {
	id := c.Param("HotelId")
	err := hotelService.DeleteHotelById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

// Maneja las solicitudes para actualizar un hotel. Obtiene los datos en un JSON,
// llama al service para realizar la actualización y envía una respuesta al client
func UpdateHotelById(c *gin.Context) {

	var hotelDto dto.HotelDto
	hotelDto.HotelId = c.Param("HotelId")
	err := c.BindJSON(&hotelDto)

	if err != nil {

		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotelDto, er := hotelService.UpdateHotelById(hotelDto)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, hotelDto)
}
