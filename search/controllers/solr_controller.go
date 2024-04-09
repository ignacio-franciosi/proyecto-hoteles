package controller

import (
	"fmt"
	"net/http"
	client "search/client"
	"search/config"
	"search/dto"
	"search/services"
	con "search/utils/connections"
	e "search/utils/errors"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var (
	Solr = services.NewSolrServiceImpl(
		(*client.SolrClient)(con.NewSolrClient(config.SOLRHOST, config.SOLRPORT, config.SOLRCOLLECTION)),
	)
)

func GetQuery(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	query := c.Param("searchQuery")

	hotelsDto, err := Solr.GetQuery(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, hotelsDto)
		return
	}

	log.Debug(hotelsDto)
	log.Debug("HOLA ACA")

	c.JSON(http.StatusOK, hotelsDto)
}

func GetQueryAllFields(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	// query := c.Param("searchQuery")

	query := "*:*"

	hotelsDto, err := Solr.GetQueryAllFields(query)
	if err != nil {
		log.Debug(err)
		c.JSON(http.StatusBadRequest, hotelsDto)
		return
	}

	c.JSON(http.StatusOK, hotelsDto)

}

func AddFromId(id string) error { // agregar e.NewBadResquest para manejar el error
	err := Solr.AddFromId(id)
	if err != nil {
		e.NewBadRequestApiError("Error adding hotel to Solr")
		return err
	}

	fmt.Println(http.StatusOK)

	return nil
}

// recibe la solicitud para eliminar un hotel, llama al servicio Solr y
// luego envía una respuesta HTTP al cliente según el resultado de la eliminación.
func Delete(id string) error {
	err := Solr.Delete(id)
	if err != nil {
		e.NewBadRequestApiError("Error deleting hotel from Solr")
		return err
	}

	fmt.Println(http.StatusOK)

	return nil
}

/*
func GetQuery(c *gin.Context) {
	// Obtener los valores de los tres campos desde la solicitud HTTP
	city := c.Query("city")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	// Construir la consulta Solr que incluye los tres campos
	query := fmt.Sprintf("city:%s AND startDate:%s AND endDate:%s", city, startDate, endDate)

	// Llamar a la función GetQuery de SolrService con la nueva consulta
	hotelsDto, err := Solr.GetQuery(query)
	if err != nil {
		log.Debug(hotelsDto)
		c.JSON(http.StatusBadRequest, hotelsDto)
		return
	}

	c.JSON(http.StatusOK, hotelsDto)
}

func AddHotelsToSolr(c *gin.Context) {
	var hotels dto.HotelsDto
	err := c.BindJSON(&hotels)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	er := Solr.AddHotelsToSolr(hotels)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

// Manejar las solicitudes de adición de hoteles al motor de búsqueda Solr.
// Extrae el ID del hotel de la solicitud HTTP, utiliza solr_services.go para agregar hoteles
func Add(c *gin.Context) {
	id := c.Param("id")
	err := Solr.Add(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	err := Solr.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, err)
}
*/
