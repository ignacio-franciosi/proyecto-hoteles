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
	"strings"

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
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/hotels/%s", config.HOTELSHOST, config.HOTELSPORT, id))
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

func (s *SolrService) GetQuery(query string) (dto.HotelsDto, e.ApiError) {
	var hotelsDto dto.HotelsDto

	// Dividir la consulta en tres partes: city, start_date y end_date
	queryParams := strings.Split(query, "_")
	if len(queryParams) != 3 {
		return hotelsDto, e.NewBadRequestApiError("Invalid query format")
	}

	city, startDate, endDate := queryParams[0], queryParams[1], queryParams[2]

	// Realizar la búsqueda en Solr utilizando los tres campos
	hotelsDto, err := s.solr.GetQuery(city, startDate, endDate)
	if err != nil {
		return hotelsDto, e.NewBadRequestApiError("Solr failed")
	}

	return hotelsDto, nil
}
