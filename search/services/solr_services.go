package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	client "search/client"
	"search/config"
	"search/dto"
	e "search/utils/errors"
	"strconv"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"
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

func (s *SolrService) GetQuery(query string) (dto.HotelsDto, e.ApiError) {
	var hotelsDto dto.HotelsDto
	queryParams := strings.Split(query, "_")

	numParams := len(queryParams)

	log.Printf("Params: %d", numParams)

	field, query := queryParams[0], queryParams[1]

	log.Printf("%s and %s", field, query)

	hotelsDto, err := s.solr.GetQuery(query, field)
	if err != nil {
		return hotelsDto, e.NewBadRequestApiError("Solr failed")
	}

	if numParams == 4 {

		startdateQuery, enddateQuery := queryParams[2], queryParams[3]
		startdateSplit := strings.Split(startdateQuery, "-")
		enddateSplit := strings.Split(enddateQuery, "-")
		startdate := fmt.Sprintf("%s%s%s", startdateSplit[0], startdateSplit[1], startdateSplit[2])
		enddate := fmt.Sprintf("%s%s%s", enddateSplit[0], enddateSplit[1], enddateSplit[2])

		sDate, _ := strconv.Atoi(startdate)
		eDate, _ := strconv.Atoi(enddate)

		log.Debug(sDate)
		log.Debug(eDate)

		// Create a channel to collect results
		resultsChan := make(chan dto.HotelDto, len(hotelsDto))

		// Create a WaitGroup
		var wg sync.WaitGroup
		var hotel dto.HotelDto

		// Iterate through each hotel and make concurrent API calls
		for _, hotel = range hotelsDto {
			wg.Add(1) // Increment the WaitGroup counter for each Goroutine
			go func(hotel dto.HotelDto) {
				defer wg.Done() // Decrement the WaitGroup counter when Goroutine is done

				// Make API call for each hotel and send the hotel ID
				result, err := s.GetHotelInfo(hotel.HotelId, sDate, eDate) // Assuming you have a function to get hotel info
				if err != nil {
					result = false
				}

				var response dto.HotelDto

				log.Debug("Adentro")
				log.Debug(result)
				log.Debug(response)

				if result == true {
					response = hotel
					resultsChan <- response
				}
			}(hotel)
		}

		// Create a slice to store the results
		var hotelResults dto.HotelsDto

		// Start a Goroutine to close the channel when all Goroutines are done
		go func() {
			wg.Wait()          // Wait for all Goroutines to finish
			close(resultsChan) // Close the channel when all Goroutines are done
		}()

		// Collect results from the channel
		for response := range resultsChan {
			hotelResults = append(hotelResults, response)
		}

		return hotelResults, nil

	}

	return hotelsDto, nil
}

func (s *SolrService) GetHotelInfo(id string, startdate int, enddate int) (bool, error) {

	resp, err := http.Get(fmt.Sprintf("http://%s:%d/hotel/availability/%s/%d/%d", config.USERAPIHOST, config.USERAPIPORT, id, startdate, enddate))

	if err != nil {
		return false, e.NewBadRequestApiError("user-res-api failed")
	}

	var body []byte
	body, _ = io.ReadAll(resp.Body)

	var responseDto dto.AvailabilityResponse
	err = json.Unmarshal(body, &responseDto)

	if err != nil {
		log.Debugf("error in unmarshal")
		return false, e.NewBadRequestApiError("getHotelInfo failed")
	}

	status := responseDto.Status
	return status, nil
}

func (s *SolrService) GetQueryAllFields(query string) (dto.HotelsDto, e.ApiError) {
	var hotelsDto dto.HotelsDto
	hotelsDto, err := s.solr.GetQueryAllFields(query)
	if err != nil {
		log.Debug(err)
		return hotelsDto, e.NewBadRequestApiError("Solr failed")
	}
	return hotelsDto, nil
}

func (s *SolrService) AddFromId(id string) e.ApiError {
	var hotelDto dto.HotelDto
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/hotels/%s", config.HOTELSHOST, config.HOTELSPORT, id))
	// resp, err := http.Get(fmt.Sprintf("http://localhost:8070/hotel/%s", id))
	if err != nil {
		log.Debugf("error getting item %s", id)
		return e.NewBadRequestApiError("error getting hotel " + id)
	}
	var body []byte
	body, _ = io.ReadAll(resp.Body)
	log.Debugf("%s", body)
	err = json.Unmarshal(body, &hotelDto)
	log.Debugf("Unmarshal result: %s", &hotelDto)
	if err != nil {
		log.Debugf("error in unmarshal of hotel %s", id)
		return e.NewBadRequestApiError("error in unmarshal of hotel")
	}
	er := s.solr.Add(hotelDto)
	log.Debug(hotelDto)
	if er != nil {
		log.Debugf("error adding to solr")
		return e.NewInternalServerApiError("Adding to Solr error", err)
	}

	return nil
}

// eliminación de un hotel del índice Solr
func (s *SolrService) Delete(id string) e.ApiError {
	err := s.solr.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

/*
func (s *SolrService) AddHotelsToSolr(hotels []dto.HotelDto) e.ApiError {
	err := s.solr.AddHotelsToSolr(hotels)
	if err != nil {
		return err
	}
	return nil
}

// agrega un hotel al motor de búsqueda Solr
// se encarga de obtener información de un hotel, deserializarla, y
// agregarla a un servidor Solr para su indexación
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

	err = s.solr.Add(hotelDto)
	if err != nil {
		logger.Debugf("error adding to solr: %v", err)
		return e.NewInternalServerApiError("Adding to Solr error", err)
	}
	return nil
}


// se encarga de tomar una cadena de consulta que contiene tres campos
// (ciudad, fecha de inicio y fecha de finalización), dividirla en partes,
// y luego realizar una búsqueda en Solr utilizando esos campos
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

*/
