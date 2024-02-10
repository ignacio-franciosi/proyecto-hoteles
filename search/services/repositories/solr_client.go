package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func (sc *SolrClient) GetQuery(query string, field string) (dto.HotelsDto, e.ApiError) {
	var response dto.SolrResponseDto
	var hotelsDto dto.HotelsDto
	q, err := http.Get(fmt.Sprintf("http://%s:%d/solr/hotelSearch/select?q=%s%s%s", config.SOLRHOST, config.SOLRPORT, field, "%3A", query))

	if err != nil {
		return hotelsDto, e.NewBadRequestApiError("error getting from solr")
	}

	defer q.Body.Close()
	err = json.NewDecoder(q.Body).Decode(&response)
	if err != nil {
		log.Printf("Response Body: %s", q.Body) // Add this line
		log.Printf("Error: %s", err.Error())
		return hotelsDto, e.NewBadRequestApiError("error in unmarshal")
	}
	hotelsDto = response.Response.Docs

	log.Printf("hotels:", hotelsDto)

	return hotelsDto, nil
}

func (sc *SolrClient) GetQueryAllFields(query string) (dto.HotelsDto, e.ApiError) {
	var response dto.SolrResponseDto
	var hotelsDto dto.HotelsDto

	q, err := http.Get(
		fmt.Sprintf("http://%s:%d/solr/hotelSearch/select?q=*:*", config.SOLRHOST, config.SOLRPORT))
	if err != nil {
		return hotelsDto, e.NewBadRequestApiError("error getting from solr")
	}

	var body []byte
	body, err = io.ReadAll(q.Body)
	if err != nil {
		return hotelsDto, e.NewBadRequestApiError("error reading body")
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return hotelsDto, e.NewBadRequestApiError("error in unmarshal")
	}

	hotelsDto = response.Response.Docs
	return hotelsDto, nil
}

func (sc *SolrClient) Add(hotelDto dto.HotelDto) e.ApiError {
	var addHotelDto dto.AddDto
	addHotelDto.Add = dto.DocDto{Doc: hotelDto}
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

func (sc *SolrClient) Delete(id string) e.ApiError {
	var deleteDto dto.DeleteDto
	deleteDto.Delete = dto.DeleteDoc{Query: fmt.Sprintf("id:%s", id)}
	data, err := json.Marshal(deleteDto)
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

/*
// serializa una busqueda de hotel a JSON y lo envia a Solr para su indexación y
// confirmar la carga para asegurarse de que los cambios se reflejen en el índice Solr
func (sc *SolrClient) Add(HotelDto dto.HotelDto) e.ApiError {
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

// toma una lista de objetos HotelDto y los convierte en formato JSON antes de enviarlos a Solr.
func (sc *SolrClient) AddHotelsToSolr(hotels []dto.HotelDto) e.ApiError {
	// Crear un slice de datos en formato JSON
	var jsonData []byte
	for _, hotel := range hotels {
		hotelData, err := json.Marshal(hotel)
		if err != nil {
			fmt.Println("Error al convertir a JSON:", err)
			return nil
		}
		jsonData = append(jsonData, hotelData...)
	}

	// URL de la colección en Solr
	solrURL := "http://localhost:8983/solr/hotels"

	// Realizar una solicitud HTTP POST para agregar los documentos a Solr
	resp, err := http.Post(solrURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error al enviar la solicitud:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Hoteles agregados exitosamente a Solr")
	} else {
		fmt.Println("Error al agregar los hoteles a Solr. Código de estado:", resp.Status)
	}

	return nil
}

func (sc *SolrClient) AddDocToSolr(HotelDto dto.HotelDto) e.ApiError {
	// Crear una instancia de tu objeto HotelDto y asignar valores
	hotel := HotelDto{
		HotelId:     "hotel123",
		Name:        "Ejemplo Hotel",
		Description: "Un gran lugar para alojarse",
		Amenities:   "Piscina, Gimnasio, Restaurante",
		Stars:       "4",
		Rooms:       100,
		Price:       150.50,
		City:        "Ejemplo Ciudad",
		Photos:      "imagen1.jpg,imagen2.jpg",
	}

	// Convertir el objeto en JSON
	jsonData, err := json.Marshal(hotel)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return nil
	}

	// URL de la colección en Solr
	solrURL := "http://localhost:8983/solr/hotels/"

	// Realizar una solicitud HTTP POST para agregar el documento a Solr
	resp, err := http.Post(solrURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error al enviar la solicitud:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Documento agregado exitosamente a Solr")
	} else {
		fmt.Println("Error al agregar el documento a Solr. Código de estado:", resp.Status)
	}
	return nil
}

// toma parámetros de búsqueda (city, startDate, endDate), realiza una solicitud a Solr,
// procesa la respuesta JSON y devuelve los resultados de la búsqueda de hoteles en el
// formato especificado en dto.HotelsDto
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


// Enviar una solicitud de eliminación a Solr para eliminar
// un documento específico basado en el ID
func (sc *SolrClient) Delete(id string) e.ApiError {

	var deleteDto dto.DeleteDto
	deleteDto.Delete = dto.DeleteDoc{Query: fmt.Sprintf("id:%s", id)}
	data, err := json.Marshal(deleteDto)
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

*/
