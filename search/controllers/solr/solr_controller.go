package solrController

import (
	"net/http"
	"search/config"
	"search/dto"
	"search/services"
	client "search/services/repositories"
	con "search/utils/solr"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

var (
	Solr = services.NewSolrServiceImpl(
		(*client.SolrClient)(con.NewSolrClient(config.SOLRHOST, config.SOLRPORT, config.SOLRCOLLECTION)),
	)
)

// Extrae la consulta de búsqueda de la solicitud HTTP, utiliza el solr_service.go para realizar la consulta
func GetQuery(c *gin.Context) {
	var hotelsArrayDto dto.HotelsArrayDto
	query := c.Param("solrQuery")

	hotelsArrayDto, err := Solr.GetQuery(query)
	if err != nil {
		log.Debug(hotelsArrayDto)
		c.JSON(http.StatusBadRequest, hotelsArrayDto)
		return
	}

	c.JSON(http.StatusOK, hotelsArrayDto)

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
