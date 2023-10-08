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
	"strings"

	log "github.com/sirupsen/logrus"
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

// envia una consulta al motor de búsqueda Solr, recibe la respuesta en formato JSON,
// decodifica en una estructura de datos Go y devuelve los hoteles encontrados
func (sc *SolrClient) GetQuery(query string) (dto.HotelsArrayDto, e.ApiError) {
	var response dto.SolrResponseDto
	var hotelsArrayDto dto.HotelsArrayDto
	query = strings.Replace(query, " ", "%20", -1)

	q, err := http.Get(fmt.Sprintf("http://%s:%d/solr/items/select?q=%s%s%s", config.SOLRHOST, config.SOLRPORT, "%3A", query))

	if err != nil {
		return hotelsArrayDto, e.NewBadRequestApiError("error getting from solr")
	}

	defer q.Body.Close()

	err = json.NewDecoder(q.Body).Decode(&response)

	if err != nil {
		log.Debug("error: ", err)
		return hotelsArrayDto, e.NewBadRequestApiError("error in unmarshal")
	}

	for _, doc := range response.Response.Docs {
		hotelArrayDto := doc
		hotelsArrayDto = append(hotelsArrayDto, hotelArrayDto)
	}

	return hotelsArrayDto, nil
}
