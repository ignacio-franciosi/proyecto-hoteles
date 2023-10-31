package solrController

import (
	"fmt"
	"net/http"
	"search/config"
	"search/services"
	client "search/services/repositories"
	con "search/utils/connections"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var (
	Solr = services.NewSolrServiceImpl(
		(*client.SolrClient)(con.NewSolrClient(config.SOLRHOST, config.SOLRPORT, config.SOLRCOLLECTION)),
	)
)

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

// recibe la solicitud para eliminar un hotel, llama al servicio Solr y
// luego envía una respuesta HTTP al cliente según el resultado de la eliminación.
func Delete(c *gin.Context) {
	id := c.Param("id")
	err := Solr.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, err)
}
