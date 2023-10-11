package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"search/config"
	"search/dto"
	e "search/utils/errors"

	logger "github.com/sirupsen/logrus"
	"github.com/stevenferrer/solr-go"
)

type SolrClient struct {
	Client     *solr.JSONClient
	Collection string
}

// serializa una busqueda de hotel a JSON y lo envia a Solr para su indexación y
// confirmar la carga para asegurarse de que los cambios se reflejen en el índice Solr
func (sc *SolrClient) AddClient(HotelDto dto.HotelDto) e.ApiError {
	var addHotelDto dto.AddDto
	addHotelDto.Add = dto.DocDto{Doc: HotelDto}
	data, err := json.Marshal(addHotelDto)

	reader := bytes.NewReader(data)
	if err != nil {
		return e.NewBadRequestApiError("Error getting json")
	}

	resp, err := sc.Client.Update(context.TODO(), sc.Collection, solr.JSON, reader)
	logger.Debug(resp)
	if err != nil {
		return e.NewBadRequestApiError("Error in solr")
	}

	er := sc.Client.Commit(context.TODO(), sc.Collection)
	if er != nil {
		logger.Debug("Error committing load")
		return e.NewInternalServerApiError("Error committing to solr", er)
	}
	return nil
}

func (s *SolrClient) GetQuery(city string, startDate string, endDate string) (dto.HotelsDto, e.ApiError) {

	var hotelsDto dto.HotelsDto

	// Construye la URL de consulta de Solr
	url := fmt.Sprintf("http://%s:%d/solr/hotels/select?q=city:%s+AND+startDate:%s+AND+endDate:%s", config.SOLRHOST, config.SOLRPORT, city, startDate, endDate)

	// Realiza una solicitud HTTP GET a la URL de Solr
	q, err := http.Get(url)

	if err != nil {
		return hotelsDto, e.NewBadRequestApiError("error getting from solr")
	}

	defer q.Body.Close()

	// Decodifica la respuesta JSON en la estructura de datos de SolrResponseDto
	var response dto.SolrResponseDto
	err = json.NewDecoder(q.Body).Decode(&response)

	if err != nil {
		return hotelsDto, e.NewBadRequestApiError("error in unmarshal")
	}

	// Asigna los documentos de la respuesta (hotels) a la variable itemsDto
	hotelsDto = response.Response.Docs

	return hotelsDto, nil
}
