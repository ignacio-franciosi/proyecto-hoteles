package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"search/config"
	"search/dto"
	client "search/services/repositories"
	e "search/utils/errors"
	"strconv"

	logger "github.com/sirupsen/logrus"
)

type SolrService struct {
	solr *client.SolrClient
}

func NewSolrServiceImpl(
	solr *client.SolrClient,
) *SolrService {
	return &SolrService{
		solr: solr,
	}
}

// agrega un hotel al motor de búsqueda Solr
func (s *SolrService) Add(id string) e.ApiError {
	var hotelDto dto.HotelDto
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/items/%s", config.HOTELSHOST, config.HOTELSPORT, id))
	if err != nil {
		logger.Debugf("error getting hotel %s: %v", id, err)
		return e.NewBadRequestApiError("error getting hotel " + id)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Debugf("bad status code %d", resp.StatusCode)
		return e.NewBadRequestApiError("bad status code " + strconv.Itoa(resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logger.Debugf("error reading response body: %v", err)
		return e.NewBadRequestApiError("error reading response body")
	}

	err = json.Unmarshal(body, &hotelDto)

	if err != nil {
		logger.Debugf("error in unmarshal of hotel %s: %v", id, err)
		return e.NewBadRequestApiError("error in unmarshal of hotel")
	}

	err = s.solr.AddClient(hotelDto)
	if err != nil {
		logger.Debugf("error adding to solr: %v", err)
		return e.NewInternalServerApiError("Adding to Solr error", err)
	}
	return nil
}

// realizar una consulta al motor de búsqueda Solr utilizando la consulta proporcionada como argumento (query)
func (s *SolrService) GetQuery(query string) (dto.HotelsArrayDto, e.ApiError) {
	var hotelsArrayDto dto.HotelsArrayDto

	hotelsArrayDto, err := s.solr.GetQuery(query)
	if err != nil {
		return hotelsArrayDto, e.NewBadRequestApiError("Solr failed")
	}

	return hotelsArrayDto, nil
}
